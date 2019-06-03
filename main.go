package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	//51.5微秒 C#
	//58.365微秒 Go

	t1 = time.Now()
	list = Suggest("internation", "en")
	elapsed = time.Since(t1)
	fmt.Println(list)
	fmt.Println("App elapsed: ", elapsed)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//certFile := "/Users/eusoft/SourceTree/sslkeys/server.pem"
	//keyFile := "/Users/eusoft/SourceTree/sslkeys/server.key"

	c, err := credentials.NewServerTLSFromFile("/Users/eusoft/SourceTree/sslkeys/server.pem", "/Users/eusoft/SourceTree/sslkeys/server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(c))
	RegisterHunspellServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	s.Serve(lis)
}
