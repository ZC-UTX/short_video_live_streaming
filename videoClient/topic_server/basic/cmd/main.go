package main

import (
	"flag"
	"fmt"
	__ "github.com/zchengutx/testproject/topics"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "topic_server/basic/init"
	"topic_server/handler/service"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterTopicServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
