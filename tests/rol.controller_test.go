package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aogallo/go-server/config"
	"github.com/aogallo/go-server/routes"

	"github.com/stretchr/testify/assert"
)

var db = make(map[string]string)

func TestPingRoute(t *testing.T) {
	database := config.ConnectDB()
	defer config.DisconnectDB(database)

	router := routes.SetupRouter(database)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
