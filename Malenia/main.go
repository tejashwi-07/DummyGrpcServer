package malenia

import (
	"context"
	"log"
	"net"
	"net/http"
	
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbMalenia "github.com/tejashwi-07/DummyGrpcServer/proto/malenia"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)


type maleniaServer struct {
	pbMalenia.UnimplementedMaleniaServiceServer
}

// MaleniaService implementation
func (*maleniaServer) HealthCheck(ctx context.Context, request *pbMalenia.HealthCheckRequest) (*pbMalenia.HealthCheckResponse, error) {
	if request.RequestValue == 1 {
		return &pbMalenia.HealthCheckResponse{HealthStatus: true}, nil
	}
	return &pbMalenia.HealthCheckResponse{HealthStatus: false}, status.Errorf(codes.Unauthenticated, "Wrong Request")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10004")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	pbMalenia.RegisterMaleniaServiceServer(s, &maleniaServer{})
	
	log.Println("Serving gRPC on 0.0.0.0:10004")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gwmux := runtime.NewServeMux()

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8092",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8092")
	log.Fatalln(gwServer.ListenAndServe())
}
