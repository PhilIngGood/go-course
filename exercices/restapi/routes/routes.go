package routes

import "github.com/gin-gonic/gin"

func CreateRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // events/1 events/4 ..
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)
}
