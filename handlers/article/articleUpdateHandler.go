package article

import (
	"net/http"

	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func UpdateArticle(c *gin.Context) {
	
	var requestData struct {
		Title       string `json:"title"`
		Body 		string `json:"body"`
		Image       string `json:"image"`
	}
	var article models.Article
	id := c.Param("id")
	if err := config.DB.Where("id=?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	// Update fields
	article.Title = requestData.Title
	article.Body = requestData.Body

	// If image is provided, update it
	if requestData.Image != "" {
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
			return
		}
		article.Image = imagePath
	}

	if err := config.DB.Save(&article).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
