package teacher

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetAllFields(c *gin.Context) {
	var listOfFields []string
	err :=config.DB.Model(&models.TeacherField{}).
	Select("name_field").
	Group("name_field").
	Scan(&listOfFields).Error
	if(err != nil) {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(listOfFields))
}