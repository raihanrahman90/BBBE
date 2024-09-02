package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	userId,_ := c.Get("userId")
	var requestData struct {
		ItemId string `json:"item_id"`
		Amount  int `json:"amount"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	cart := models.Cart{
		ItemID: requestData.ItemId,
		Amount: requestData.Amount,
		UserID: userId.(string),
	}

	if err := config.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Success"))
}
