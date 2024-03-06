package main

import (
	"database/sql"
	"net"

	"github.com/Math2121/go-rpc/internal/database"
	"github.com/Math2121/go-rpc/internal/pb"
	"github.com/Math2121/go-rpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
