package transaction

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	userId, _ := c.Get("userId")
	type requestItem struct {
		Id		string	`json:"id"`
		Amount	int		`json:"amount"`
		Price	int		`json:"price"`
	}
	var requestData []requestItem;

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var orderItems []models.OrderItem
	for _, data := range requestData {
		orderItem := models.OrderItem{
			ItemID: data.Id,
			Amount: data.Amount,
			Price: data.Price,
		}
		orderItems = append(orderItems, orderItem)		
	}

	order := models.Order{
		UserID: userId.(string),
		Date: time.Now(),
		OrderItem: orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(order)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
