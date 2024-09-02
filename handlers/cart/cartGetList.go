package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListCart(c *gin.Context) {
	var cart []models.Cart
	query := config.DB.Model(&models.Cart{})
	offset, limit, page := utils.GetPagination(c)

	var totalItems int64
	query.Count(&totalItems)
	if err := query.Offset(offset).Limit(limit).Preload("Item").Find(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := responseList(cart)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit, page))

}
