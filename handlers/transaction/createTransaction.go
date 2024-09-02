package transaction

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	userId, _ := c.Get("userId")
	var requestItem struct {
		Item		[]string	`json:"item"`
		Address		string	  `json:"address"`
		City		string	  `json:"city"`
		Province	string		`json:"province"`
	}
	var listCart []models.Cart

	if err := c.ShouldBindJSON(&requestItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("id in ?", requestItem.Item).Find(&listCart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to get cart"))
		return
	}
	var orderItems []models.OrderItem
	for _, cart := range listCart {
		orderItem := models.OrderItem{
			ItemID: cart.ItemID,
			Amount: cart.Amount,
			Price:  cart.Item.Price,
			Image:  cart.Item.Image,
			Name:   cart.Item.Name,
		}
		orderItems = append(orderItems, orderItem)
	}

	order := models.Order{
		UserID:    userId.(string),
		Date:      time.Now(),
		OrderItem: orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Delete(&listCart)
	// Response
	responseData := responseDetailTransaction(order)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
