package handler

import (
	"assignment/internal/app/model"
	service "assignment/internal/app/service"
	"assignment/pkg/article/validator"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ArticleGroup - handles request related to article
type ArticleGroup struct {
	articleService *service.ArticleService
	logger         *log.Logger
}

//SaveArticleHandler - to save article data into database
func (ag *ArticleGroup) SaveArticleHandler(c *gin.Context) {
	articleObj := model.Article{}
	err := c.ShouldBind(&articleObj)
	if err != nil {
		ag.logger.Println("USER_PARAMETER_BIND_ERROR : ", err)
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if ok, err := validator.ValidateInputs(&articleObj); !ok {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  http.StatusUnprocessableEntity,
			"message": err,
			"data":    nil,
		})
		return
	}
	ID, err := ag.articleService.SaveArticleService(c.Request.Context(), articleObj)
	if err != nil {
		ag.logger.Println(err)
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Success",
		"data": map[string]interface{}{
			"id": ID,
		},
	})
	return
}

// FetchArticleByIDHandler - to get article object by ID from database
func (ag *ArticleGroup) FetchArticleByIDHandler(c *gin.Context) {
	ID := c.Param("article_id")
	var articleView model.Article
	err := ag.articleService.FetchArticleByIDService(c.Request.Context(), ID, &articleView)
	if err != nil {
		ag.logger.Println("failed to fetch article: ", err)
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "failed to fetch article",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    articleView,
	})
}

//FetchArticlesHandler - to get articles from database
func (ag *ArticleGroup) FetchArticlesHandler(c *gin.Context) {
	articleViews := []model.Article{}
	err := ag.articleService.FetchArticlesService(c.Request.Context(), &articleViews)
	if err != nil {
		ag.logger.Println("failed to fetch articles: ", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "failed to fetch articles",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    articleViews,
	})
}
