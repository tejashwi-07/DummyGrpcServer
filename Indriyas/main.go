package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbIndriyas "github.com/tejashwi-07/DummyGrpcServer/proto/indriyas"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type indriyasServer struct {
	pbIndriyas.UnimplementedIndriyasServiceServer
}

// IndpbIndriyasService implementation
func (*indriyasServer) HealthCheck(ctx context.Context, request *pbIndriyas.HealthCheckRequest) (*pbIndriyas.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbIndriyas.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbIndriyas.HealthCheckResponse{HealthStatus: false}, status.Errorf(codes.Unauthenticated, "Wrong Request")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	pbIndriyas.RegisterIndriyasServiceServer(s, &indriyasServer{})

	log.Println("Serving gRPC on 0.0.0.0:10002")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gwmux := runtime.NewServeMux()

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8091")
	log.Fatalln(gwServer.ListenAndServe())
}
