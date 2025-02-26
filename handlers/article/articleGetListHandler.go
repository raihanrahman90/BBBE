package article

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	var article []models.Article
	if err := config.DB.Find(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}
	responseData := responseList(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
	return ;
}