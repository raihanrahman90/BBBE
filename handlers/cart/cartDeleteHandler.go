package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCart(c *gin.Context) {
	id := c.Param("id")
	var cart models.Cart
	if err := config.DB.Where("itemId = ? and userId = ?", id).First(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	if err := config.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(cart))

}
