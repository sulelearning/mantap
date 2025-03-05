package api

import (
	"fmt"

	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	"github.com/Zulhaidir/microservice/mantap/token"
	"github.com/Zulhaidir/microservice/mantap/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

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

	/* Daftarkan validator dengan Gin */
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// routing untuk users
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// membuat group untuk authentication route
	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	// routing untuk accounts
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.PUT("/accounts/:id", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	// routing untuk transfers
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

/* Start runs the HTTP server on a specific address */
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
