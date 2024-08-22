package article

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var requestData struct {
		Title       string `json:"title"`
		Body 		string `json:"body"`
		Image       string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	var existRecord models.Article
	if err := config.DB.Where("title=?", requestData.Title).First(&existRecord).Error; err==nil{
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Title already exist"))
		return
	}

	// Save base64 image to disk
	imagePath, err := utils.SaveBase64Image(requestData.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
		return
	}

	// Create teacher object
	article := models.Article{
		Title:       requestData.Title,
		Image:       imagePath,
		Body: 		 requestData.Body,
	}

	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Internal Server Error"))
		return
	}

	responseData := response(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
