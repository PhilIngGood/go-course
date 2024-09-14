package main

import (
	"course.go/restapi/db"
	"course.go/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.CreateRoutes(server)

	server.Run("127.0.0.1:8080")

}
