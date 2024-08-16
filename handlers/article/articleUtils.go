package article

import (
	"rumahbelajar/models"
	"rumahbelajar/utils"
)

type responseDTO struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Body	string	`json:"body"`
	Image	string	`json:"image"`	
}

func response(data models.Article) responseDTO{
	var response responseDTO
	response.ID = data.ID
	response.Body = data.Body
	response.Title = data.Title
	response.Image = utils.GetImageUrl(data.Image)
	return response
}

func responseList(datas []models.Article) []responseDTO {
	var results []responseDTO
	for _, data := range datas {
		result:=response(data)
		results = append(results, result)
	}
	return results
}