package main

import (
	"context"
	"flag"
	"time"

	pb "github.com/douglarek/grpc-gateway-demo/proto/gen/go/echo/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var target = flag.String("target", "127.0.0.1:9090", "grpc server address and port")
var timeout = flag.Int("timeout", 10, "request time second")

func main() {

	flag.Parse()

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	conn, err := grpc.DialContext(ctx, *target, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewEchoServiceClient(conn)
	t := time.NewTimer(time.Duration(*timeout) * time.Second)
	q := make(chan struct{})
	wait := make(chan struct{})

	go func() {
		for {
			select {
			case <-q:
				close(wait)
				return
			default:
				grpclog.Infoln(c.Echo(context.Background(), &pb.StringMessage{Value: "world"}))
			}
		}
	}()

	<-t.C
	close(q)

	<-wait
}
