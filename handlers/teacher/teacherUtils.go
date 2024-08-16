package teacher

import (
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"
)

type responseType struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Fields      []string `json:"fields"`
}

func responseList(datas []models.Teacher) []responseType {
	var results []responseType
	for _, data := range datas {
		result:=response(data)
		results = append(results, result)
	}
	return results
}

func response(data models.Teacher) responseType{
	var fieldNames []string
	for _, fields := range data.Fields {
		fieldNames = append(fieldNames, fields.NameField)
	}
	result := responseType{
		Id:          data.ID,
		Name:        data.Name,
		Fields:      fieldNames,
		Image:       utils.GetImageUrl(data.Image),
		Description: data.Description,
	}
	utils.LogObject(result)
	utils.LogObject(data)
	return result;
}

func generateFields(fields []string) []models.TeacherField {
	fieldsOfStudy := make([]models.TeacherField, len(fields))

	for i, name := range fields{
		fieldsOfStudy[i] = models.TeacherField{NameField: name}
	}
	return fieldsOfStudy;
}

func updateAllFields(id string, field []string) []models.TeacherField{
	deleteAllFields(id);
	fields := generateFields(field);
	return fields;
}

func deleteAllFields(id string) {
	var fields []models.TeacherField
	if err := config.DB.Where("teacher_id = ?", id).Find(&fields).Error; err != nil {
		return;
	}
	config.DB.Delete(&fields)
}