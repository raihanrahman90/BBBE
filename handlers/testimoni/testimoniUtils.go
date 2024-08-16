package testimoni

import (
	"rumahbelajar/models"
	"rumahbelajar/utils"
)

type responseType struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Testimoni 	string   `json:"testimoni"`
	Image       string   `json:"image"`
}

func responseList(datas []models.Testimoni) []responseType {
	var results []responseType
	for _, data := range datas {
		result:=response(data)
		results = append(results, result)
	}
	return results
}

func response(data models.Testimoni) responseType{
	result := responseType{
		Id:          data.ID,
		Name:        data.Name,
		Testimoni:	 data.Testimoni,
		Image:       utils.GetImageUrl(data.Image),
	}
	utils.LogObject(result)
	utils.LogObject(data)
	return result;
}