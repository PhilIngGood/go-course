package routes

import (
	"course.go/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(server *gin.Engine) {

	authMiddleware := server.Group("/")
	authMiddleware.Use(middlewares.Authenticate)
	authMiddleware.POST("/events", createEvents)
	authMiddleware.PUT("/events/:id", updateEvent)
	authMiddleware.DELETE("/events/:id", deleteEvent)
	authMiddleware.POST("/events/:id/register", registerForEvent)
	authMiddleware.DELETE("/events/:id/register", cancelRegistration)
	authMiddleware.GET("/registration", getRegistration)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1 events/4 ..
	server.POST("/signup", signup)
	server.POST("/login", login)
}
