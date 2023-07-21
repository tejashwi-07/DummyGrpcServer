package timesquared

import (
	"context"
	"log"
	"net"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbTimeSquared "github.com/tejashwi-07/DummyGrpcServer/proto/timesquared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type timesquaredServer struct {
	pbTimeSquared.UnimplementedTimeSquaredServiceServer
}

// MaleniaService implementation
func (*timesquaredServer) HealthCheck(ctx context.Context, request *pbTimeSquared.HealthCheckRequest) (*pbTimeSquared.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbTimeSquared.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbTimeSquared.HealthCheckResponse{HealthStatus: false}, status.Errorf(codes.Unauthenticated, "Wrong Request")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10005")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	pbTimeSquared.RegisterTimeSquaredServiceServer(s, &timesquaredServer{})
	
	log.Println("Serving gRPC on 0.0.0.0:10005")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gwmux := runtime.NewServeMux()

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8094",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8094")
	log.Fatalln(gwServer.ListenAndServe())
}
