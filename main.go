package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbHelloWorld "github.com/tejashwi-07/DummyGrpcServer/proto/helloworld"
	pbApexDrive "github.com/tejashwi-07/DummyGrpcServer/proto/apexdrive"
	pbMalenia "github.com/tejashwi-07/DummyGrpcServer/proto/malenia"
	pbTimeSquared "github.com/tejashwi-07/DummyGrpcServer/proto/timesquared"
	pbIndriyas "github.com/tejashwi-07/DummyGrpcServer/proto/indriyas"
	pbNeith "github.com/tejashwi-07/DummyGrpcServer/proto/neith"
	pbGateway "github.com/tejashwi-07/DummyGrpcServer/proto/gateway"
)

// Server struct representing our service implementation
type server struct{}
type apexDriveServer struct{
	pbApexDrive.UnimplementedApexDriveServiceServer
}
type maleniaServer struct{
	pbMalenia.UnimplementedMaleniaServiceServer
}
type timeSquaredServer struct{
	pbTimeSquared.UnimplementedTimeSquaredServiceServer
}
type indriyasServer struct{
	pbIndriyas.UnimplementedIndriyasServiceServer
}
type neithServer struct{
	pbNeith.UnimplementedNeithServiceServer
}
type gatewayServer struct{
	pbGateway.UnimplementedGatewayServiceServer
}

// SayHello is the implementation of the SayHello method defined in the proto file
func (*server) SayHello(_ context.Context, in *pbHelloWorld.HelloRequest) (*pbHelloWorld.HelloReply, error) {
	return &pbHelloWorld.HelloReply{Message: in.Name + " world"}, nil
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

func (*gatewayServer) ApexDriveStart(ctx context.Context, request *pbGateway.ApexDriveStartRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Apexdrive service started.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) MaleniaStart(ctx context.Context, request *pbGateway.MaleniaStartRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Malenia service started.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) TimeSquaredStart(ctx context.Context, request *pbGateway.TimeSquaredStartRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("TimeSquared service started.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) IndriyasStart(ctx context.Context, request *pbGateway.IndriyasStartRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Indriyas service started.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) NeithStart(ctx context.Context, request *pbGateway.NeithStartRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Neith service started.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) ApexDriveStop(ctx context.Context, request *pbGateway.ApexDriveStopRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Apexdrive service stopped.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) MaleniaStop(ctx context.Context, request *pbGateway.MaleniaStopRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Malenia service stopped.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) TimeSquaredStop(ctx context.Context, request *pbGateway.TimeSquaredStopRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("TimeSquared service stopped.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) IndriyasStop(ctx context.Context, request *pbGateway.IndriyasStopRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Indriyas service stopped.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) NeithStop(ctx context.Context, request *pbGateway.NeithStopRequest) (*pbGateway.GatewayResponse, error) {
	fmt.Println("Neith service stopped.")
	return &pbGateway.GatewayResponse{}, nil
}

func (*gatewayServer) ApexDriveStatus(ctx context.Context, request *pbGateway.ApexDriveStatusRequest) (*pbGateway.GatewayResponseforStatus, error) {
	if request.StatusValue == 1 {
		return &pbGateway.GatewayResponseforStatus{ServiceStatus: true}, nil
	}
	return &pbGateway.GatewayResponseforStatus{ServiceStatus: false}, nil
}

func (*gatewayServer) MaleniaStatus(ctx context.Context, request *pbGateway.MaleniaStatusRequest) (*pbGateway.GatewayResponseforStatus, error) {
	if request.StatusValue == 1 {
		return &pbGateway.GatewayResponseforStatus{ServiceStatus: true}, nil
	}
	return &pbGateway.GatewayResponseforStatus{ServiceStatus: false}, nil
}

func (*gatewayServer) TimeSquaredStatus(ctx context.Context, request *pbGateway.TimeSquaredStatusRequest) (*pbGateway.GatewayResponseforStatus, error) {
	if request.StatusValue == 1 {
		return &pbGateway.GatewayResponseforStatus{ServiceStatus: true}, nil
	}
	return &pbGateway.GatewayResponseforStatus{ServiceStatus: false}, nil
}

func (*gatewayServer) IndriyasStatus(ctx context.Context, request *pbGateway.IndriyasStatusRequest) (*pbGateway.GatewayResponseforStatus, error) {
	if request.StatusValue == 1 {
		return &pbGateway.GatewayResponseforStatus{ServiceStatus: true}, nil
	}
	return &pbGateway.GatewayResponseforStatus{ServiceStatus: false}, nil
}

func (*gatewayServer) NeithStatus(ctx context.Context, request *pbGateway.NeithStatusRequest) (*pbGateway.GatewayResponseforStatus, error) {
	if request.StatusValue == 1 {
		return &pbGateway.GatewayResponseforStatus{ServiceStatus: true}, nil
	}
	return &pbGateway.GatewayResponseforStatus{ServiceStatus: false}, nil
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
	pbHelloWorld.RegisterGreeterServer(s, &server{})
	pbApexDrive.RegisterApexDriveServiceServer(s, &apexDriveServer{})
	pbMalenia.RegisterMaleniaServiceServer(s, &maleniaServer{})
	pbTimeSquared.RegisterTimeSquaredServiceServer(s, &timeSquaredServer{})
	pbIndriyas.RegisterIndriyasServiceServer(s, &indriyasServer{})
	pbNeith.RegisterNeithServiceServer(s, &neithServer{})
	pbGateway.RegisterGatewayServiceServer(s, &gatewayServer{})


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
	// Register the Greeter service with the gRPC-Gateway
	err = pbHelloWorld.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
