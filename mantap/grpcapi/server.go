package grpcapi

import (
	"fmt"

	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	"github.com/Zulhaidir/microservice/mantap/pb"
	"github.com/Zulhaidir/microservice/mantap/token"
	"github.com/Zulhaidir/microservice/mantap/util"
)

// Server serves gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer create a new server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker terdiri dari JWTMaker dan PasetoMaker, diantara keduanya PasetoMaker yang lebih baik, jadi saya menggunakannya
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
