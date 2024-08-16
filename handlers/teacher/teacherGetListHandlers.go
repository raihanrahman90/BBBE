package teacher

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)


func GetTeacher(c *gin.Context) {
	var teachers []models.Teacher
	if err := config.DB.Preload("Fields").Find(&teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to get data"))
		return
	}
	response := responseList(teachers)

	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}