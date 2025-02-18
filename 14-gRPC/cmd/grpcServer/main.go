package main

import (
	"database/sql"
	"net"

	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
	"github.com/devfullcycle/14-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Abre conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Cria banco de dados
	categoryDb := database.NewCategory(db)
	// Cria o serviço que manipula o banco de dados
	categoryService := service.NewCategoryService(*categoryDb)

	// Cria servidor gRPC
	grpcServer := grpc.NewServer()
	// Attacha o serviço ao servidor gRPC
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	// Habilita interação com o Evans
	reflection.Register(grpcServer)

	// Abre porta TCP pra acessar o serviço gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
