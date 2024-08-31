package item

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemById(c *gin.Context) {
	id := c.Param("id")

	var item models.Item

	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := response(item)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
