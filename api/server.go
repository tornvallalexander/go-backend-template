package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"net/http"
)

// Server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server with routing
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.DELETE("/users/:username", server.deleteUser)

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