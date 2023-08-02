package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbApexDrive "github.com/tejashwi-07/DummyGrpcServer/proto/apexdrive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server struct representing our service implementation

type apexdriveServer struct {
	pbApexDrive.UnimplementedApexDriveServiceServer
}

func (*apexdriveServer) HealthCheck(ctx context.Context, request *pbApexDrive.HealthCheckRequest) (*pbApexDrive.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbApexDrive.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbApexDrive.HealthCheckResponse{HealthStatus: false}, status.Errorf(codes.Unauthenticated, "Wrong Request")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	pbApexDrive.RegisterApexDriveServiceServer(s, &apexdriveServer{})

	reflection.Register(s)

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:10001")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

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
