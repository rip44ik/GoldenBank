package api

import (
	"github.com/gin-gonic/gin"
	db "gitlab.com/xfx1/goldbank/db/sqlc"
)

// этот сервер будет обрабатывать все запросы банка
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New server creates and new HTTP server setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
