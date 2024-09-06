package transaction

import (
	"bbbe/config"
	"bbbe/enums"
	"bbbe/models"
	"bbbe/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	userId, _ := c.Get("userId")
	var request map[string]int
	var listCart []models.Cart

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemIds := getAllItemId(request)
	var items []models.Item
	if err := config.DB.Where("id in ?", itemIds).Find(&items).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return
	}

	if err := config.DB.Where("user_id = ?", userId).Find(&listCart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return
	}
	total := 0;
	var orderItems []models.OrderItem
	for _, item := range items {
		orderItem := models.OrderItem{
			ItemID: item.ID,
			Amount: request[item.ID],
			Name:   item.Name,
			Price:  item.Price,
			Image:  item.Image,
		}
		total = total + (orderItem.Amount * orderItem.Price)
		orderItems = append(orderItems, orderItem)
	}

	order := models.Order{
		UserID:    userId.(string),
		Date:      time.Now(),
		OrderItem: orderItems,
		Total: 		total,
		Status: 	string(enums.WAITING_PAYMENT),
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

func getAllItemId(data map[string]int) []string{
    keys := make([]string, 0, len(data))
    for k := range data {
        keys = append(keys, k)
    }
	return keys;
}