package main

import (
	"context"
	"log"
	"net"

	pb "github.com/douglarek/grpc-gateway-demo/proto/gen/go/echo/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
