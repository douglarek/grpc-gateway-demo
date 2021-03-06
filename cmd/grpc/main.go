package main

import (
	"context"
	"log"
	"net"

	pb "github.com/douglarek/grpc-gateway-demo/proto/gen/go/echo/service/v1"
	hb "github.com/douglarek/grpc-gateway-demo/proto/gen/go/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":9090"
)

// server is used to implement v1.EchoServiceServer.
type server struct {
	pb.UnimplementedEchoServiceServer
}

// Echo implements v1.EchoServiceServer
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	grpclog.Infof("Received: %v", in.GetValue())
	return &pb.StringMessage{Value: "Hello " + in.GetValue()}, nil
}

// healthServer implements GRPC Health Checking Protocol
type healthServer struct {
	hb.UnimplementedHealthServer
}

func (s *healthServer) Check(ctx context.Context, in *hb.HealthCheckRequest) (*hb.HealthCheckResponse, error) {
	return &hb.HealthCheckResponse{Status: hb.HealthCheckResponse_SERVING}, nil
}

func (s *healthServer) Watch(in *hb.HealthCheckRequest, _ hb.Health_WatchServer) error {
	// Example of how to register both methods but only implement the Check method.
	return status.Error(codes.Unimplemented, "unimplemented")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterEchoServiceServer(s, &server{})
	hb.RegisterHealthServer(s, &healthServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
