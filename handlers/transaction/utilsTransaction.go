package transaction

import (
	"bbbe/models"
	"time"
)

type responseTransactionDTO struct {
	ID        string                     `json:"id"`
	ItemName  string                     `json:"item_name"`
	ItemPrice int                        `json:"item_price"`
	ItemImage string                     `json:"item_image"`
	CountItem int                        `json:"count_item"`
	Total     int                        `json:"total"`
	Date      time.Time                  `json:"date"`
	Status    string                     `json:"status"`
	Item      responseTransactionItemDTO `json:"item"`
}

type responseTransactionItemDTO struct {
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemImage string `json:"item_image"`
	Amount    int    `json:"amount"`
	SubTotal  int    `json:"sub_total"`
}

func responseTransaction(data models.Order) responseTransactionDTO {
	var response responseTransactionDTO
	var item models.OrderItem
	item = data.OrderItem[0]
	response.ID = data.ID
	response.ItemName = item.Name
	response.ItemPrice = item.Price
	response.ItemImage = item.Image
	response.Status = data.Status
	response.Date = data.Date
	response.CountItem = len(data.OrderItem)
	return response
}

func responseListTransaction(datas []models.Order) []responseTransactionDTO {
	var results []responseTransactionDTO
	for _, data := range datas {
		result := responseTransaction(data)
		results = append(results, result)
	}
	return results

}

func responseDetailTransaction(data models.Order) responseTransactionDTO {
	var response responseTransactionDTO
	var item models.OrderItem
	item = data.OrderItem[0]
	response.ID = data.ID
	response.ItemName = item.Name
	response.ItemPrice = item.Price
	response.ItemImage = item.Image
	response.Status = data.Status
	response.Date = data.Date
	response.CountItem = len(data.OrderItem)
	return response
}
