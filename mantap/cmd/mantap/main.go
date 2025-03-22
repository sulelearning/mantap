package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/Zulhaidir/microservice/mantap/api"
	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	_ "github.com/Zulhaidir/microservice/mantap/docs/statik"
	"github.com/Zulhaidir/microservice/mantap/grpcapi"
	"github.com/Zulhaidir/microservice/mantap/pb"
	"github.com/Zulhaidir/microservice/mantap/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	/* Load Config */
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Tidak dapat load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Msg("Tidak dapat konek ke database")
	}

	// run DB migration, disini kita menjalankan migrations tanpa perlu menggunakan pada command line: "make migrate-up"
	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)

	// runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Msg("tidak dapat membuat instance migrate baru")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msg("gagal untuk menjalankan migrate up")
	}

	log.Info().Msg("db migrate suceesfully")
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("Tidak dapat membuat server")
	}

	// logger gRPC API
	grpcLogger := grpc.UnaryInterceptor(grpcapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)

	// opsional: menadaftarkan reflection service untuk server grpc
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msg("Tidak dapat membuat listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("tidak dapat menjalankan server gRPC")
	}
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("Tidak dapat membuat server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("Tidak dapat Menghandle register server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg("tidak dapat membuat statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("Tidak dapat membuat listener")
	}

	log.Info().Msgf("start HTTP Gateway server at %s", listener.Addr().String())
	handler := grpcapi.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Msg("tidak dapat menjalankan server HTTP Gateway")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("Tidak dapat membuat server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("Tidak dapat menjalankan server")
	}
}
