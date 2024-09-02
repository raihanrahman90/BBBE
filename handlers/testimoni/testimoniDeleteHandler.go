package testimoni

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTestimoni(c *gin.Context) {
	id := c.Param("id")
	var testimoni models.Testimoni
	if err := config.DB.Where("id = ?", id).First(&testimoni).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return
	}

	if err := config.DB.Delete(&testimoni).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return;
	}

	responseData := response(testimoni)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	
}