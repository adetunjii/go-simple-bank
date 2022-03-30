package api

import (
	. "github.com/Adetunjii/simplebank/db/repository"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store IStore
	router *gin.Engine
}

func CreateNewServer(store IStore) *Server {
	server := &Server {store: store}
	router := gin.Default()

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)

	//TRANSFERS
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) StartServer(addr string)error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}