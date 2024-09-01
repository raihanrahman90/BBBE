package item

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	var item []models.Item
	query := config.DB.Model(&models.Item{})

	if minPrice := c.Query("minPrice"); minPrice != "" {
		query = query.Where("price > ?", minPrice)
	}
	if maxPrice := c.Query("maxPrice"); maxPrice != "" {
		query = query.Where("price < ?", maxPrice)
	}
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	sortBy, sortOrder := utils.GetSorting(c)
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	offset, limit, page := utils.GetPagination(c)

	if err := query.Offset(offset).Limit(limit).Find(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	var totalItems int64
	query.Count(&totalItems)

	responseData := responseList(item)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit, page))
}
