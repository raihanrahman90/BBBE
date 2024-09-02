package address

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListAddress(c *gin.Context) {
	var address []models.Address
	userId,_ := c.Get("userId")
	query := config.DB.Model(&models.Address{})
	offset, limit, page := utils.GetPagination(c)
	query = query.Where("userId = ?", userId)

	var totalItems int64
	query.Count(&totalItems)
	if err := query.Offset(offset).Limit(limit).Find(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponsePagination(address, int(totalItems), limit, page))

}
