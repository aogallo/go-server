package main

import (
	"net/http"

	"github.com/aogallo/go-server/config"
	"github.com/aogallo/go-server/models"
	"github.com/aogallo/go-server/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = make(map[string]string)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// User Routes
	routes.SetupUserRoutes(r, db)

	// Rol Routes
	routes.SetupRolRoutes(r, db)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	//
	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}
	//
	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	database := config.ConnectDB()
	defer config.DisconnectDB(database)

	// Migrate the schemas
	database.AutoMigrate(&models.User{}, &models.Rol{})

	// // validations
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterStructValidation(controllers.RolStructLevelValidation, models.Rol{})
	// }

	r := SetupRouter(database)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
