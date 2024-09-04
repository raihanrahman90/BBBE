package cart

import (
	"bbbe/models"
	"bbbe/utils"
)

type responseDTO struct {
	ItemID	  string `json:"itemId"`
	ItemName  string `json:"itemName"`
	ItemPrice int    `json:"itemPrice"`
	ItemImage string `json:"itemImage"`
}

func response(data models.Cart) responseDTO {
	var response responseDTO
	item := data.Item
	response.ItemID = item.ID
	response.ItemName = item.Name
	response.ItemPrice = item.Price
	response.ItemImage = utils.GetImageUrl(item.Image)
	return response
}

func responseList(datas []models.Cart) []responseDTO {
	var results []responseDTO
	for _, data := range datas {
		result := response(data)
		results = append(results, result)
	}
	return results
}
