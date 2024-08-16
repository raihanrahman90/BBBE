package testimoni

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)


func GetTestimoni(c *gin.Context) {
	var Testimonis []models.Testimoni
	if err := config.DB.Find(&Testimonis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to get data"))
		return
	}
	response := responseList(Testimonis)

	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}