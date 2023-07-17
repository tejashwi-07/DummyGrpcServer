package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbApexDrive "github.com/tejashwi-07/DummyGrpcServer/proto/apexdrive"
	pbAuth "github.com/tejashwi-07/DummyGrpcServer/proto/auth"
	pbDocker "github.com/tejashwi-07/DummyGrpcServer/proto/docker"
	pbIndriyas "github.com/tejashwi-07/DummyGrpcServer/proto/indriyas"
	pbMalenia "github.com/tejashwi-07/DummyGrpcServer/proto/malenia"
	pbNeith "github.com/tejashwi-07/DummyGrpcServer/proto/neith"
	pbTimeSquared "github.com/tejashwi-07/DummyGrpcServer/proto/timesquared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Server struct representing our service implementation

type apexDriveServer struct {
	pbApexDrive.UnimplementedApexDriveServiceServer
}
type maleniaServer struct {
	pbMalenia.UnimplementedMaleniaServiceServer
}
type timeSquaredServer struct {
	pbTimeSquared.UnimplementedTimeSquaredServiceServer
}
type indriyasServer struct {
	pbIndriyas.UnimplementedIndriyasServiceServer
}
type neithServer struct {
	pbNeith.UnimplementedNeithServiceServer
}
type dockerServer struct {
	pbDocker.UnimplementedDockerServiceServer
}

type authServer struct {
	pbAuth.UnimplementedAuthServiceServer
}

func (s *authServer) Authenticate(ctx context.Context, req *pbAuth.AuthRequest) (*pbAuth.AuthResponse, error) {
	// Retrieve the product key from the request
	productKey := req.ProductKey

	// Perform the authentication logic
	// Replace this with your actual authentication implementation

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

func (*apexDriveServer) Healthcheck(ctx context.Context, request *pbApexDrive.HealthCheckRequest) (*pbApexDrive.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbApexDrive.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbApexDrive.HealthCheckResponse{HealthStatus: false}, nil
}

// MaleniaService implementation
func (*maleniaServer) HealthCheck(ctx context.Context, request *pbMalenia.HealthCheckRequest) (*pbMalenia.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbMalenia.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbMalenia.HealthCheckResponse{HealthStatus: false}, nil
}

// TimeSquaredService implementation
func (*timeSquaredServer) HealthCheck(ctx context.Context, request *pbTimeSquared.HealthCheckRequest) (*pbTimeSquared.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbTimeSquared.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbTimeSquared.HealthCheckResponse{HealthStatus: false}, nil
}

func (*indriyasServer) Healthcheck(ctx context.Context, request *pbIndriyas.HealthCheckRequest) (*pbIndriyas.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbIndriyas.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbIndriyas.HealthCheckResponse{HealthStatus: false}, nil
}

func (*neithServer) HealthCheck(ctx context.Context, request *pbNeith.HealthCheckRequest) (*pbNeith.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbNeith.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbNeith.HealthCheckResponse{HealthStatus: false}, nil
}

func (*dockerServer) StartService(ctx context.Context, request *pbDocker.DockerRequest) (*pbDocker.DockerResponse, error) {
	fmt.Printf("%v started.", request.ServiceName)
	return &pbDocker.DockerResponse{}, nil
}

func (*dockerServer) StopService(ctx context.Context, request *pbDocker.DockerRequest) (*pbDocker.DockerResponse, error) {
	fmt.Printf("%v Stopped.", request.ServiceName)
	return &pbDocker.DockerResponse{}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pbAuth.RegisterAuthServiceServer(s, &authServer{})
	pbApexDrive.RegisterApexDriveServiceServer(s, &apexDriveServer{})
	pbMalenia.RegisterMaleniaServiceServer(s, &maleniaServer{})
	pbTimeSquared.RegisterTimeSquaredServiceServer(s, &timeSquaredServer{})
	pbIndriyas.RegisterIndriyasServiceServer(s, &indriyasServer{})
	pbNeith.RegisterNeithServiceServer(s, &neithServer{})
	pbDocker.RegisterDockerServiceServer(s, &dockerServer{})

	s = withAuthInterceptor(s)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Create a new ServeMux for the gRPC-Gateway
	gwmux := runtime.NewServeMux()

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

func withAuthInterceptor(server *grpc.Server) *grpc.Server {
	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(authInterceptor),
	}
	server = grpc.NewServer(serverOpts...)
	return server
}

// authInterceptor is the authentication interceptor function.
func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Perform the authentication check here
	if err := authenticate(ctx); err != nil {
		// Authentication failed, return an error
		return nil, status.Errorf(codes.Unauthenticated, "Authentication failed: %v", err)
	}

	// Authentication succeeded, continue with the request
	return handler(ctx, req)
}

// authenticate performs the authentication check.
// Replace this with your actual authentication logic.
func authenticate(ctx context.Context) error {
	// Retrieve the metadata from the gRPC request context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "Metadata not found")
	}

	// Extract the authentication token from the metadata
	token := md.Get("authorization")
	if len(token) == 0 {
		return status.Error(codes.Unauthenticated, "Missing authentication token")
	}

	// Validate the authentication token
	// Replace this with your actual token validation logic

	if len(token) != 10 {
		return status.Error(codes.Unauthenticated, "Invalid authentication token")
	}

	return nil
}
