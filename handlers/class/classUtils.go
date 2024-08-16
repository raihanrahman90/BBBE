package class

import (
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"
)

type responseDTO struct {
	ID			string	`json:"id"`
	Title       string 	`json:"title"`
	Price       int    	`json:"price"`
	Description string 	`json:"description"`
	Image       string 	`json:"image"`
	Module 		[]string`json:"modules"`
}

func response(data models.Class) responseDTO{
	var fieldNames []string
	for _, fields := range data.Modules {
		fieldNames = append(fieldNames, fields.Module)
	}
	var response responseDTO
	response.ID = data.ID
	response.Title = data.Title
	response.Price = data.Price
	response.Description = data.Description
	response.Image = utils.GetImageUrl(data.Image)
	response.Module = fieldNames
	return response
}

func responseList(datas []models.Class) []responseDTO {
	var results []responseDTO
	for _, data := range datas {
		result:=response(data)
		results = append(results, result)
	}
	return results
}


func generateModules(modulesName []string) []models.ClassModules {
	modules := make([]models.ClassModules, len(modulesName))

	for i, name := range modulesName{
		modules[i] = models.ClassModules{Module: name}
	}
	return modules;
}

func updateAllModules(id string, field []string) []models.ClassModules{
	deleteAllModules(id);
	modules := generateModules(field);
	return modules;
}

func deleteAllModules(id string) {
	var modules []models.ClassModules
	if err := config.DB.Where("class_id = ?", id).Find(&modules).Error; err != nil {
		return;
	}
	config.DB.Delete(&modules)
}