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
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) StemWord(ctx context.Context, in *HunspellRequest) (*HunspellReply, error) {
	return &HunspellReply{WordList:StemWord("allons", "fr")}, nil
}

func (s *server) Suggestion(ctx context.Context, in *HunspellRequest) (*HunspellReply, error) {
	return &HunspellReply{WordList:StemWord("allons", "fr")}, nil
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
	RegisterHunspellServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
