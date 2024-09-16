package transaction

import (
	"bbbe/models"
	"bbbe/utils"
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
	ProofOfPayment string				 `json:"proof_of_payment"`
	Username	string						`json:"username"`
	City		string						`json:"city"`
	Province		string
	Address		string
	Item	  []responseTransactionItemDTO	`json:"item"`
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
	response.ItemImage = utils.GetImageUrl(item.Image)
	response.Status = data.Status
	response.Date = data.Date
	response.CountItem = len(data.OrderItem)
	response.Total = data.Total
	response.ProofOfPayment = utils.GetImageUrl(data.ProofOfPayment)
	response.Username = data.User.Username
	response.City = data.City
	response.Address = data.Address
	response.Province = data.Province
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
	var responseItem []responseTransactionItemDTO
	item := data.OrderItem[0]
	response.ID = data.ID
	response.ItemName = item.Name
	response.ItemPrice = item.Price
	response.ItemImage = utils.GetImageUrl(item.Image)
	response.Status = data.Status
	response.Date = data.Date
	response.CountItem = len(data.OrderItem)
	response.Total = data.Total
	response.Username = data.User.Username
	response.ProofOfPayment = utils.GetImageUrl(data.ProofOfPayment)
	if data.ProofOfPayment != "" {
		response.ProofOfPayment = utils.GetImageUrl(data.ProofOfPayment)
	}
	for _, dataOrder := range data.OrderItem{
		item := responseItemTransaction(dataOrder)
		responseItem = append(responseItem, item)
	}
	response.Item = responseItem;
	return response
}

func responseItemTransaction(data models.OrderItem) responseTransactionItemDTO{
	var response responseTransactionItemDTO
	response.Amount = data.Amount
	response.ItemImage = utils.GetImageUrl(data.Image)
	response.ItemName = data.Name
	response.ItemPrice = data.Price
	response.SubTotal = data.Amount * data.Price
	return response;
}
