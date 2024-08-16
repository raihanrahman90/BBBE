package landingpage

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

type landingPageRequest struct {
	Key 	string `json:"key"`
	Type 	string `json:"type"`
	Value 	string `json:"value"`
	Image	string `json:"image"`
}

func UpdateDataLandingPage(c *gin.Context) {
		var landingPage models.LandingPage
		var requestData landingPageRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
			return
		}
		config.DB.Where("key=?", requestData.Key).First(&landingPage)
		landingPage.Key = requestData.Key;
		landingPage.Type = requestData.Type;
		if(requestData.Type != "image") {
			landingPage.Value = requestData.Value;
		}else{
			updateWithImage(&requestData, &landingPage, c)
		}

		if err := config.DB.Save(&landingPage).Error; err != nil{
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save data"))
			return;
		}

		c.JSON(http.StatusOK, landingPage)
}

func updateWithImage(requestData *landingPageRequest, landingPage *models.LandingPage, c *gin.Context) {
	if requestData.Image != "" {
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed save new Image"))
			return
		}
		landingPage.Value = imagePath
	}
}