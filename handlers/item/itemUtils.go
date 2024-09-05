package item

import (
	"bbbe/models"
	"bbbe/utils"
)

type responseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category	*string `json:"category"`
	Rating		int		`json:"rating"`
}

func response(data models.Item) responseDTO {
	var response responseDTO
	response.ID = data.ID
	response.Name = data.Name
	response.Price = data.Price
	response.Description = data.Description
	response.Category = data.Category
	response.Image = utils.GetImageUrl(data.Image)
	response.Rating = 4
	return response
}

func responseList(datas []models.Item) []responseDTO {
	var results []responseDTO
	for _, data := range datas {
		result := response(data)
		results = append(results, result)
	}
	return results
}
