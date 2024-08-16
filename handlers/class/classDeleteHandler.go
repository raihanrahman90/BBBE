package class

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func DeleteClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	if err := config.DB.Where("id = ?", id).First(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return
	}

	deleteAllModules(id);

	if err := config.DB.Delete(&class).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"));
		return;
	}

	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	
}