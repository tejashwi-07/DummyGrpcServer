package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbNeith "github.com/tejashwi-07/DummyGrpcServer/proto/neith"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

type neithServer struct {
	pbNeith.UnimplementedNeithServiceServer
}

// MaleniaService implementation
func (*neithServer) HealthCheck(ctx context.Context, request *pbNeith.HealthCheckRequest) (*pbNeith.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbNeith.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbNeith.HealthCheckResponse{HealthStatus: false}, status.Errorf(codes.Unauthenticated, "Wrong Request")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10003")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	pbNeith.RegisterNeithServiceServer(s, &neithServer{})

	log.Println("Serving gRPC on 0.0.0.0:10003")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gwmux := runtime.NewServeMux()

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8093",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8093")
	log.Fatalln(gwServer.ListenAndServe())
}
