package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"github.com/tornvallalexander/go-backend-template/utils"
	"net/http"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	config utils.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server with routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.DELETE("/users/:username", server.deleteUser)
	router.GET("/users/", server.listUsers)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"errors": err.Error()}
}

func checkErr(err error) (status int, res gin.H) {
	switch err {
	case sql.ErrNoRows:
		return http.StatusNotFound, errorResponse(err)
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation", "foreign_key_violation":
			return http.StatusForbidden, errorResponse(pqErr)
		}
	}

	return http.StatusInternalServerError, errorResponse(err)
}
