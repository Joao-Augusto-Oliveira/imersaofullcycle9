package server

import (
	"github.com/Joao-Augusto-Oliveira/imersaofullcycle9/infrastructure/grpc/pb"
	"github.com/Joao-Augusto-Oliveira/imersaofullcycle9/infrastructure/grpc/service"
	"github.com/Joao-Augusto-Oliveira/imersaofullcycle9/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp","0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen tpc port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}