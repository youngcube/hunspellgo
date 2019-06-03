package main

import (
	"log"
	"net"

	"./spellcheck"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50052"

type server struct{}

func (s *server) StemWord(ctx context.Context, req *spellcheck.StemRequest) (*spellcheck.WordListReply, error) {
	return &spellcheck.WordListReply{WordList: spellcheck.StemWord(req.Word, req.Lang)}, nil
}

func (s *server) StemLine(ctx context.Context, req *spellcheck.StemRequest) (*spellcheck.StemLineReply, error) {
	return &spellcheck.StemLineReply{Result: spellcheck.StemLine(req.Word, req.Lang)}, nil
}

func (s *server) Suggestion(ctx context.Context, req *spellcheck.SuggestionRequest) (*spellcheck.WordListReply, error) {
	return &spellcheck.WordListReply{WordList: spellcheck.Suggest(req.Word, req.Lang, req.Count)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	spellcheck.RegisterHunspellServer(s, &server{})
	log.Println("start listen", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
