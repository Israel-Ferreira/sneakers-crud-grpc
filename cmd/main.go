package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Israel-Ferreira/sneakers-crud-grpc/internal/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/Israel-Ferreira/sneakers-crud-grpc/pkg/pb/proto"
)

func main() {
	fmt.Println("Sneaker Crud GRPC")

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterSneakerServiceServer(server, &svc.SneakerService{})

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalln("Erro ao subir o servidor gRPC: ", err.Error())
	}
}
