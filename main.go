package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/db"
	"github.com/sinclare210/Backend.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
