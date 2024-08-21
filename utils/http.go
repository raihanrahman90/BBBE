package utils

import (
	"net/http"
	"strconv"

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
	TotalPage	int			`json:"total_page"`
}

func SuccessResponse(response interface{}) DefaultResponse{
	var defaultResponse DefaultResponse
	defaultResponse.Data 		= response
	defaultResponse.IsSuccess 	= true
	defaultResponse.Status		= 200
	defaultResponse.Message		= "success"
	return defaultResponse;
}

func SuccessResponsePagination(response interface{}, totalItem int, pageSize int) DefaultResponse{
	var defaultResponse DefaultResponse
	defaultResponse.Data 		= response
	defaultResponse.IsSuccess 	= true
	defaultResponse.Status		= 200
	defaultResponse.Message		= "success"
	defaultResponse.TotalPage	= (totalItem + pageSize -1)/(pageSize)
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

func GetPagination(c *gin.Context) (int, int){
	page := c.DefaultQuery("page", "1")
    pageSize := c.DefaultQuery("pageSize", "10")
	var pageInt, pageSizeInt int
    var err error
	if pageInt, err = strconv.Atoi(page); err != nil || pageInt < 1 {
        pageInt = 1
    }

    if pageSizeInt, err = strconv.Atoi(pageSize); err != nil || pageSizeInt < 1 {
        pageSizeInt = 10
    }

	offset := (pageInt - 1) * pageSizeInt
	return offset, pageSizeInt;
}