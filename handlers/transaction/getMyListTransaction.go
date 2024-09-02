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

	userId, _ := c.Get("userId")
	query = query.Where("user_id = ?", userId);

	if status := c.Query("status"); status != "" {
		query = query.Where("LOWER(status) = LOWER(?)", status)
	}

	sortBy, sortOrder := utils.GetSorting(c)
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	offset, limit, page := utils.GetPagination(c)

	var totalItems int64
	query.Count(&totalItems)
	if err := query.Offset(offset).Limit(limit).Find(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := responseListTransaction(order)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit, page))

}
