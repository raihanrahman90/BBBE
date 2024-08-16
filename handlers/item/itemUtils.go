package item

import (
	"rumahbelajar/models"
	"rumahbelajar/utils"
)

type responseDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
}

func response(data models.Item) responseDTO {
	var fieldNames []string
	for _, fields := range data.ItemImage {
		var path = utils.GetImageUrl(fields.Path)
		fieldNames = append(fieldNames, path)
	}
	var response responseDTO
	response.ID = data.ID
	response.Name = data.Name
	response.Price = data.Price
	response.Description = data.Description
	response.Image = fieldNames
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
