package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/models"
)

func GetEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events, try again later"})
		return
	}
	context.JSON(http.StatusOK,events)
}

func CreateEvent(context *gin.Context){
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse resquest"})
		fmt.Println(err)
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create event, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func GetEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
		if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not parse event id"})
		return
	}
	context.JSON(http.StatusOK, event)
}

