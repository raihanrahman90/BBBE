package cart

import (
	"bbbe/models"
	"bbbe/utils"
)

type responseDTO struct {
	ID          string   `json:"id"`
	ItemName    string   `json:"item_name"`
	ItemPrice   int      `json:"item_price"`
	ItemImage 	string	 `json:"item_image"`
	Amount		int			`json:"amount"`
	SubTotal	int			`json:"sub_total"`

}

func response(data models.Cart) responseDTO {
	var response responseDTO
	item := data.Item;
	response.ID = data.ID
	response.ItemName	= item.Name;
	response.ItemPrice	=	item.Price;
	if item.ItemImage != nil {
		response.ItemImage	=	utils.GetImageUrl(item.ItemImage[0].Path)
	}
	response.Amount = data.Amount;
	response.SubTotal = item.Price*data.Amount;
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
