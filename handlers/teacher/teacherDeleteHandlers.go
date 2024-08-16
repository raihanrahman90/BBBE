package teacher

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := config.DB.Where("id = ?", id).First(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"))
		return
	}
	deleteAllFields(id);
	if err:= config.DB.Delete(&teacher).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return;
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(true));
	return;
	
}