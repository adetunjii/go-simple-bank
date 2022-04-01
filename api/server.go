package api

import (
	"fmt"
	. "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/token"
	"github.com/Adetunjii/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config       util.Config
	tokenFactory token.TokenFactory
	store        IStore
	router       *gin.Engine
}

func CreateNewServer(config util.Config, store IStore) (*Server, error) {
	tokenFactory, err := token.NewJwtFactory(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token factory: %v", err)
	}

	server := &Server{store: store, tokenFactory: tokenFactory, config: config}

	// Create custom validator
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) StartServer(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/login", server.loginUser)
	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)

	//TRANSFERS
	router.POST("/transfers", server.createTransfer)

	server.router = router
}
