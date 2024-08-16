package teacher

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)


func GetTeacherById(c *gin.Context) {
	id := c.Param("id")

	var teacher models.Teacher

	if err := config.DB.Where("id = ?", id).Preload("Fields").First(&teacher).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data not found"))
		return
	}
	response := response(teacher)

	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}