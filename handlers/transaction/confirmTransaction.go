package transaction

import (
	"bbbe/config"
	"bbbe/enums"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfirmTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Order

	var requestData struct{
		ReceiptNumber	string	`json:"receipt_number"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	if err:= config.DB.Where("id = ?", id).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return;
	}

	transaction.Status = string(enums.VALIDATED_PAYMENT);
	transaction.ReceiptNumber = requestData.ReceiptNumber;
	if err := config.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to update data"))
		return;
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(transaction)); 
}