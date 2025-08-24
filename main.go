package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/db"
	"github.com/sinclare210/Backend.git/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // Get,put,post,patch,delete
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events,err := models.GetAllEvents()
		if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"msessage": "Could not parse request data"})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create events,n try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Events created!", "event": event})

}
