package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Backend.git/utils"
)

func Authenticate(context *gin.Context){

	token := context.Request.Header.Get("Authorization")
	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not authorised"})
		return 
	}

	user_id,err := utils.VerifyToken(token)
	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not authorized"})
		return 
	}
	context.Set("user_id",user_id)
	context.Next()
	
}