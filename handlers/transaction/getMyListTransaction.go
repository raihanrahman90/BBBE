package transaction

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyListTransaction(c *gin.Context) {
	var order []models.Order
	query := config.DB.Model(&models.Order{})

	if status := c.Query("status"); status != "" {
		query = query.Where("LOWER(status) = LOWER(?)", status)
	}

	sortBy, sortOrder := utils.GetSorting(c)
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	userId, _ := c.Get("userId")
	query = query.Where("user_id = ?", userId);

	offset, limit, page := utils.GetPagination(c)

	if status := c.Query("status"); status != "" {
		query = query.Where("LOWER(status) = LOWER(?)", status)
	}

	var totalItems int64
	query.Count(&totalItems)
	if err := query.Offset(offset).Limit(limit).Preload("OrderItem").Preload("User").Order("date DESC").Find(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := responseListTransaction(order)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit, page))

}
