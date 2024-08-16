package item

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func CreateItemImage(c *gin.Context) {
	var requestData struct {
		ItemId string `json:"item_id"`
		Image  string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	imagePath, err := utils.SaveBase64Image(requestData.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
		return
	}

	itemImage := models.ItemImage{
		Path:   imagePath,
		ItemID: requestData.ItemId,
	}

	if err := config.DB.Create(&itemImage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Success"))
}
