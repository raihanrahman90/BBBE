package landingpage

import (
	"bbbe/models"
	"bbbe/utils"
)

type responseType struct {
	Id          string   `json:"id"`
	Key        	string   `json:"key"`
	Type 		string   `json:"type"`
	Value		string	 `json:"value"`
	Image		string	 `json:"image"`
}

func responseList(datas []models.LandingPage) map[string]string {
	mapLanding := make(map[string]string)
	for _, data := range datas {
		if data.Type == "image" {
			mapLanding[data.Key] = utils.GetImageUrl(data.Value)
		}else{
			mapLanding[data.Key] = data.Value
		}
	}
	return mapLanding
}