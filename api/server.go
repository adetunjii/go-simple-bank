package api

import (
	. "github.com/Adetunjii/simplebank/db/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *Store
	router *gin.Engine
}

func CreateNewServer(store *Store) *Server {
	server := &Server {store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) StartServer(addr string)error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}