package user

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []models.User

	query := config.DB.Model(&models.User{})

	sortBy, sortOrder := utils.GetSorting(c)
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	offset, limit, page := utils.GetPagination(c)

	if err := query.Offset(offset).Limit(limit).Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	var totalItems int64
	query.Count(&totalItems)

	c.JSON(http.StatusOK, utils.SuccessResponsePagination(user, int(totalItems), limit, page))
}