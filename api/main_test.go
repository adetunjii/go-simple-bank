package api

import (
	db "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func NewTestServer(t *testing.T, store db.IStore) *Server {
	config := util.Config{
		SecretKey:     util.RandomString(32),
		TokenDuration: time.Minute,
	}

	server, err := CreateNewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
