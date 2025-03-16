package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Zulhaidir/microservice/mantap/api"
	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	"github.com/Zulhaidir/microservice/mantap/grpcapi"
	"github.com/Zulhaidir/microservice/mantap/pb"
	"github.com/Zulhaidir/microservice/mantap/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	/* Load Config */
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Tidak dapat konek ke database:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
	runGinServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Tidak dapat membuat server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	// opsional: menadaftarkan reflection service untuk server grpc
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Tidak dapat membuat listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("tidak dapat menjalankan server gRPC")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Tidak dapat membuat server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Tidak dapat menjalankan server:", err)
	}
}
