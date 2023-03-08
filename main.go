package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	hellopb "github.com/FeLvi-zzz/sample_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	hellopb.UnimplementedHelloServer
}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("rcv")
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339Nano)

	return &hellopb.HelloResponse{
		Hostname: hostname,
		Time:     now,
	}, nil
}

func (*server) StreamHello(req *hellopb.HelloRequest, stream hellopb.Hello_StreamHelloServer) error {
	fmt.Println("rcv")
	for {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}

		now := time.Now().Format(time.RFC3339Nano)

		stream.Send(&hellopb.HelloResponse{
			Hostname: hostname,
			Time:     now,
		})

		time.Sleep(time.Second * 1)
	}
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)

	hellopb.RegisterHelloServer(s, &server{})

	reflection.Register(s)

	log.Printf("%s", l.Addr())
	log.Println("--------------")

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
