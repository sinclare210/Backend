package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents) //GET,POST,PUT,PATCH,DELETE
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	server.POST("/signup",signUp)
	server.POST("login",login)
}