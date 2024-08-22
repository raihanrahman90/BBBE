package testimoni

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

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