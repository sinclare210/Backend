package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/models"
	
)

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events, try again later"})
		return
	}
	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse resquest"})
		fmt.Println(err)
		return
	}

	user_id := context.GetInt64("user_id")

	
	event.UserID = user_id
	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create event, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context){
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

func updateEvent(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event id"})
		return
	}
	user_id := context.GetInt64("user_id")
	event, err := models.GetEventById(eventId)
		if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fecth the event"})
		return
	}
	
	

	if event.UserID != user_id{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not athorized to update event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse resquest"})
		
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
		if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not update the event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Event updated successfully"})

}

func deleteEvent(context *gin.Context){
	eventId,err := strconv.ParseInt(context.Param("id"),10,64)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event id"})
		return
	}

	deleteEvent, err := models.GetEventById(eventId)
		if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fecth the event"})
		return
	}
//comparism
	user_id := context.GetInt64("user_id")
	if deleteEvent.UserID != user_id{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not athorized to delete event"})
		return
	}

	

	err = deleteEvent.Delete()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully!"})



}