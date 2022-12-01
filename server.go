package main

import (
	"log"
	"net"

	"github.com/Aziz0310/bootcamp/grpc_example/proto-gen/dice"
	dService "github.com/Aziz0310/bootcamp/grpc_example/services/dice"
	"google.golang.org/grpc"
)

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, &dService.TutorialService{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
