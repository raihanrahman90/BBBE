package transaction

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListTransaction(c *gin.Context) {
	var order []models.Order
	query := config.DB.Model(&models.Order{})

	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	sortBy, sortOrder := utils.GetSorting(c)
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	offset, limit := utils.GetPagination(c)

	if err := query.Offset(offset).Limit(limit).Find(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	var totalItems int64
	query.Count(&totalItems)

	responseData := responseListTransaction(order)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit))

}
