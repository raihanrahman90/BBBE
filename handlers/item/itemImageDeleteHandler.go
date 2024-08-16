package item

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func DeleteItemImage(c *gin.Context) {
	id := c.Param("id")
	var itemImage models.ItemImage
	if err := config.DB.Where("id = ?", id).First(&itemImage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"))
		return
	}

	if err := config.DB.Delete(&itemImage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(itemImage))

}
