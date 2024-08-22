package testimoni

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func DeleteTestimoni(c *gin.Context) {
	id := c.Param("id")
	var testimoni models.Testimoni
	if err := config.DB.Where("id = ?", id).First(&testimoni).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return
	}

	if err := config.DB.Delete(&testimoni).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return;
	}

	responseData := response(testimoni)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	
}