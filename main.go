package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	//"./binding"
	//"fmt"
	"golang.org/x/net/context"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main()  {

	t1 := time.Now()
	list := StemWord("allons", "fr")
	elapsed := time.Since(t1)
	fmt.Println(list)
	fmt.Println("App elapsed: ", elapsed)

	t1 = time.Now()
	list = Suggest("internation", "en")
	elapsed = time.Since(t1)
	fmt.Println(list)
	fmt.Println("App elapsed: ", elapsed)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
