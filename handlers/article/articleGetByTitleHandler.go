package article

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func GetArticleByTitle(c *gin.Context) {
	title := c.Param("title")

	var article models.Article

	if err := config.DB.Where("title = ?", title).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := response(article)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}