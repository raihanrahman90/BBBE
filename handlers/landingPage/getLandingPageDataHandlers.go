package landingpage

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func GetLandingPageData(c *gin.Context) {
	var landingPage []models.LandingPage
	if err := config.DB.Find(&landingPage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to get Landing Page Data"))
		return
	}
	response:=responseList(landingPage)
	c.JSON(http.StatusOK, utils.SuccessResponse(response));
}