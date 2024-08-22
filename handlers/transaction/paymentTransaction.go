package transaction

import (
	"bbbe/config"
	"bbbe/enums"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaymentTransaction(c *gin.Context) {
	id := c.Param("id")
	var requestData struct {
		Image        string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var transaction models.Order

	if err:= config.DB.Where("id = ?", id).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return;
	}

	imagePath, err := utils.SaveBase64Image(requestData.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
		return
	}

	transaction.Status = string(enums.WAITING_VALIDATION);
	transaction.ProofOfPayment = imagePath
	if err := config.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to update data"))
		return;
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(transaction));
}