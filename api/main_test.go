package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"github.com/tornvallalexander/go-backend-template/utils"
	"os"
	"testing"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
