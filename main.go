package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/db"
	"github.com/sinclare210/Backend.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()		

	server.GET("/events", routes.GetEvents) //GET,POST,PUT,PATCH,DELETE
	server.GET("/events/:id", routes.GetEvent)
	server.POST("/events", routes.CreateEvent)
	server.Run(":8080")//localhost:8080
}

