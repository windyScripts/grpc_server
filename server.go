package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server)Add(ctx context.Context, req *pb.AddRequest)(*pb.AddResponse, error){
	return &pb.AddResponse{
		Sum: req.A + req.B,
	},nil
}

func main() {

	// IMPORT CERTIFICATES. WIll have to create them first.

	cert:= "cert.pem"
	key:= "key.pem"

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err!= nil {
		log.Fatalln("Failed to load credentials", err)
	}

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	// ADD CREDENTIALS HERE
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})

	// skipped

	fmt.Println("Server started on 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
