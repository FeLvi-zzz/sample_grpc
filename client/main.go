package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	hellopb "github.com/FeLvi-zzz/sample_grpc/proto"
)

func main() {
	flag.Parse()
	target := flag.Arg(0)
	fmt.Println(target)

	conn, err := grpc.Dial(
		target,
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				// Time:                10 * time.Second,
				// Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := hellopb.NewHelloClient(conn)

	for {
		res, err := c.Hello(context.TODO(), &hellopb.HelloRequest{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", res)

		time.Sleep(5 * time.Second)
	}
}
