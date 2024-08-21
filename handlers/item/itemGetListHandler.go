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
	query := config.DB.Model(&models.Item{})

	if minPrice := c.Query("minPrice"); minPrice != "" {
		query = query.Where("price > ?", minPrice)
	}
	if maxPrice := c.Query("maxPrice"); maxPrice != "" {
		query = query.Where("price < ?", maxPrice);
	}
	if name := c.Query("name"); name != "" {
        query = query.Where("name LIKE ?", "%"+name+"%")
    }
	
	offset, limit := utils.GetPagination(c)
	if err := query.Offset(offset).Limit(limit).Preload("ItemImage").Find(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	var totalItems int64;
	query.Count(&totalItems);

	responseData := responseList(item)
	c.JSON(http.StatusOK, utils.SuccessResponsePagination(responseData, int(totalItems), limit))
}
