package tests

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/aogallo/go-server/internal/db"
	"github.com/aogallo/go-server/internal/routes"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	path := filepath.Join("../", ".env.test")
	database := db.ConnectDB(path)
	defer db.DisconnectDB(database)

	router := routes.SetupRouter(database)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
