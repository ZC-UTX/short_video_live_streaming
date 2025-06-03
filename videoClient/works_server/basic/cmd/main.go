package main

import (
	"flag"
	"fmt"
	__ "github.com/zchengutx/testproject/works"
	"google.golang.org/grpc"
	"log"
	"net"
<<<<<<< HEAD
=======
	_ "works_server/basic/init"

>>>>>>> 8a09d8c643121d1a9daf3be3d46ad033437c593c
	"works_server/handler/service"
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
	__.RegisterWorksServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
