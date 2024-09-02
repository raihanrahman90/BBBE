package item

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var class models.Item
	if err := config.DB.Where("id = ?", id).First(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	if err := config.DB.Delete(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))

}
