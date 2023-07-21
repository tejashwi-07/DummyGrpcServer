package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/dgrijalva/jwt-go"
	pbEngine "github.com/tejashwi-07/DummyGrpcServer/proto/engine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type engineServer struct {
	pbEngine.UnimplementedMicroserviceControllerServer
}

func (s *engineServer) Authenticate(ctx context.Context, req *pbAuth.AuthRequest) (*pbAuth.AuthResponse, error) {
	productKey := req.ProductKey

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
	response := &pbAuth.AuthResponse{
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

func (s *engineServer) StartServer(ctx context.Context, req *pb.ServerRequest) (*pb.ServerResponse, error) {
	server := req.ServiceName
	// Here, you can use the Docker SDK to start the corresponding microservice container.
	// Implement the logic to start the microservice container based on the 'microservice' parameter.

	// Connect to the Docker daemon
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Failed to connect to Docker daemon: %v", err)
		return &pb.ServerResponse{
			Message: "Failed to connect to Docker daemon",
		}, err
	}
	defer cli.Close()

	// Set the image name for the ApexDrive microservice
	imageName := server + "-image:latest" // Replace with the actual image name

	// Pull the latest image (optional, but recommended)
	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Failed to pull Docker image: %v", err)
		return &pb.ServerResponse{
			Message: "Failed to pull Docker image",
		}, err
	}
	defer reader.Close()

	containerName := server + "-container"
	// Create and start the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, containerName)
	if err != nil {
		log.Printf("Failed to create Docker container: %v", err)
		return &pb.ServerResponse{
			Message: "Failed to create Docker container",
		}, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Failed to start Docker container: %v", err)
		return &pb.ServerResponse{
			Message: "Failed to start Docker container",
		}, err
	}

	// Container started successfully
	return &pb.ServerResponse{
		Message: service + " started successfully",
	}, nil

}

func (s *engineServer) StopServer(ctx context.Context, req *pb.ServerRequest) (*pb.ServerResponse, error) {
	server := req.ServiceName
	// Here, you can use the Docker SDK to start the corresponding microservice container.
	// Implement the logic to start the microservice container based on the 'microservice' parameter.
	// Connect to the Docker daemon
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("Failed to connect to Docker daemon: %v", err)
		return &pb.ServerResponse{
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
		return &pb.ServerResponse{
			Message: "Failed to list containers",
		}, err
	}

	// If the container is not found, it means it's not running
	if len(runningContainers) == 0 {
		log.Printf("%s container is not running.", microservice)
		return &pb.ServerResponse{
			Message: server + " container is not running",
		}, nil
	}

	// Stop the container
	containerID := runningContainers[0].ID
	timeout := 5 * time.Second
	if err := cli.ContainerStop(ctx, containerID, &timeout); err != nil {
		log.Printf("Failed to stop Docker container: %v", err)
		return &pb.ServerResponse{
			Message: "Failed to stop Docker container",
		}, err
	}
	// Container stopped successfully
	return &pb.ServerResponse{
		Message: server + " stopped successfully",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pbEngine.RegisterMicroserviceControllerServer(grpcServer, &engineServer{})

	log.Println("Serving gRPC on 0.0.0.0:10000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
