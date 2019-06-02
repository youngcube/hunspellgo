package main

import (
	"fmt"
	"time"

	//"./binding"
	//"fmt"
	//"golang.org/x/net/context"
	//"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	//"io/ioutil"
	//"log"
	//"net"
	//"os"
	//"time"
	//"log"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	log.Printf("Received: %v", in.Name)
//	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
//}

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

	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &server{})
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
