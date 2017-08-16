package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	pb "github.com/alextanhongpin/go-grpc-gateway/hello"
)

type helloServer struct {
}

func (s *helloServer) EchoPost(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println(msg)
	return msg, nil
}
func (s *helloServer) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg, ctx)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}
