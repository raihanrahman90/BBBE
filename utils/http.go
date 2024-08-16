package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBody(requestData interface{}, c *gin.Context) interface{}{
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return requestData;
}

type DefaultResponse struct {
	IsSuccess	bool		`json:"isSuccess"`
	Status 		int			`json:"status"`
	Data 		interface{}	`json:"data"`
	Message		string		`json:"message"`
}

func SuccessResponse(response interface{}) DefaultResponse{
	var defaultResponse DefaultResponse
	defaultResponse.Data 		= response
	defaultResponse.IsSuccess 	= true
	defaultResponse.Status		= 200
	defaultResponse.Message		= "success"
	return defaultResponse;
}

func FailedResponse(message string) DefaultResponse{
		var defaultResponse DefaultResponse
		defaultResponse.Data 		= nil
		defaultResponse.IsSuccess 	= false
		defaultResponse.Status		= 400
		defaultResponse.Message		= message
		return defaultResponse;
}