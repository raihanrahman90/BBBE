package item

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	var item []models.Item
	if err := config.DB.Preload("ItemImage").Find(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := responseList(item)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
