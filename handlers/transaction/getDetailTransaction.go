package transaction

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetailTranscation(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.Where("id = ?", id).Preload("OrderItem").Preload("User").Find(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to get data"))
		return
	}

	response := responseDetailTransaction(order)
	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}
