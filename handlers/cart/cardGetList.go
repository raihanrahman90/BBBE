package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListCard(c *gin.Context) {
	var cart []models.Cart
	query := config.DB.Model(&models.Cart{})
	offset, limit := utils.GetPagination(c)
	
	if err := query.Offset(offset).Limit(limit).Preload("Item").Preload("Item.ItemImage").Find(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	var totalItems int64;
	query.Count(&totalItems);

	responseData := responseList(cart)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit))
	
}