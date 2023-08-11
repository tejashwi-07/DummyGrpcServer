package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	pbEngine "github.com/Tarran-Sidhaarth/DummyGrpcServer/proto/engine"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type engineServer struct {
	pbEngine.UnimplementedMicroserviceControllerServer
}

func (s *engineServer) Authenticate(ctx context.Context, req *pbEngine.AuthRequest) (*pbEngine.AuthResponse, error) {
	productKey := req.ProductKey
	log.Printf("product key: %v", productKey)
	if len(productKey) != 10 {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":   "1234567890",
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour).Unix(), // Expiration time
	}

	// Generate the JWT token with the claims and secret key
	secretKey := "my-secret-key"

	// Generate the authentication token
	token, err := GenerateJWTToken(claims, secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	// Create the response with the generated token
	response := &pbEngine.AuthResponse{
		TokenValue: token,
	}

	return response, nil
}

func GenerateJWTToken(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *engineServer) StartServer(ctx context.Context, req *pbEngine.ServerRequest) (*pbEngine.ServerResponse, error) {
	server := req.ServiceName
	log.Printf("Trying to ignite service : %s\n", server)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Failed to connect to Docker daemon: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to connect to Docker daemon",
		}, err
	}
	defer cli.Close()

	// Set the image name for the ApexDrive microservice
	imageName := server + "-image:latest" // Replace with the actual image name
	log.Printf("image:%v", imageName)

	containerName := server + "-container"
	// Create and start the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"10001/tcp": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"10001/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "10001",
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		log.Printf("Failed to create Docker container: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to create Docker container",
		}, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Failed to start Docker container: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to start Docker container",
		}, err
	}
	go GetStats(context.Background(), resp.ID)
	// Container started successfully
	return &pbEngine.ServerResponse{
		Message: server + " started successfully-----------",
	}, nil

}

func (s *engineServer) StopServer(ctx context.Context, req *pbEngine.ServerRequest) (*pbEngine.ServerResponse, error) {
	server := req.ServiceName
	// Here, you can use the Docker SDK to start the corresponding microservice container.
	// Implement the logic to start the microservice container based on the 'microservice' parameter.
	// Connect to the Docker daemon

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Failed to connect to Docker daemon: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to connect to Docker daemon",
		}, err
	}
	defer cli.Close()

	// Find the container by name
	containerName := server + "-container" // Replace with the appropriate naming convention
	runningContainers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("name", containerName)),
	})
	if err != nil {
		log.Printf("Failed to list containers: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to list containers",
		}, err
	}

	// If the container is not found, it means it's not running
	if len(runningContainers) == 0 {
		log.Printf("%s container is not running.", server)
		return &pbEngine.ServerResponse{
			Message: server + " container is not running",
		}, nil
	}

	// Stop the container
	containerID := runningContainers[0].ID
	timeout := 5 * time.Second
	timeoutSeconds := int(timeout.Seconds())
	stopOptions := container.StopOptions{
		Timeout: &timeoutSeconds,
	}

	if err := cli.ContainerStop(ctx, containerID, stopOptions); err != nil {
		log.Printf("Failed to stop Docker container: %v", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to stop Docker container",
		}, err
	}
	// Removing the container

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := cli.ContainerRemove(ctx, containerID, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return &pbEngine.ServerResponse{
			Message: "Failed to stop Docker container",
		}, err
	}

	//Container stopped and removed successfully
	return &pbEngine.ServerResponse{
		Message: server + " stopped successfully",
	}, nil

}

func GetStats(ctx context.Context, containerId string) {
	server := "req.ServiceName"
	log.Printf("Getting stats of the service %s\n", server)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to connect to Docker daemon: %v", err)
	}
	defer cli.Close()
	stats, err := cli.ContainerStats(ctx, "apexdrive-container", true)
	if err != nil {
		log.Fatalf("Failed to get stats of Docker container: %v", err)
	}
	defer stats.Body.Close()
	var stat types.StatsJSON
	decoder := json.NewDecoder(stats.Body)
	for {
		if err := decoder.Decode(&stat); err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("Failed to get stats of Docker container : %v", err)
		}
		log.Printf("CPU Usage: %d\n", stat.CPUStats.CPUUsage.TotalUsage)
		log.Printf("Memory Usage: %d\n", stat.MemoryStats.Usage)
	}
}

func (s *engineServer) GetServerStats(req *pbEngine.ServerRequest, stream pbEngine.MicroserviceController_GetServerStatsServer) error {
	server := req.ServiceName
	containerName := server + "-container"
	log.Printf("Getting stats of the service %s\n", server)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Failed to connect to Docker daemon: %v", err)
		return err
	}
	defer cli.Close()
	ctx := context.Background()
	stats, err := cli.ContainerStats(ctx, containerName, true)
	if err != nil {
		log.Printf("Failed to get stats of Docker container: %v", err)
		return err
	}
	defer stats.Body.Close()
	var stat types.StatsJSON
	decoder := json.NewDecoder(stats.Body)
	for {
		if err := decoder.Decode(&stat); err != nil {
			if err == io.EOF {
				return err
			}
			log.Fatalf("Failed to get stats of Docker container : %v", err)
		}
		res := &pbEngine.ServerResponse{
			Message: "Cpu Usage " + strconv.FormatUint(stat.CPUStats.CPUUsage.TotalUsage, 10) + " Memory Usage " + strconv.FormatUint(stat.MemoryStats.Usage, 10),
		}
		if err := stream.Send(res); err != nil {
			log.Printf("%v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pbEngine.RegisterMicroserviceControllerServer(grpcServer, &engineServer{})
	reflection.Register(grpcServer)
	log.Println("Serving gRPC on 0.0.0.0:10000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
