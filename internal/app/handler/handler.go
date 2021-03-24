package handler

import (
	"assignment/internal/app/repo"
	"assignment/internal/app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//New - init the Event Router
func New(conn *mongo.Database, logger *log.Logger) (http.Handler, error) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	articleRepo, err := repo.NewArticleRepo(conn)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	articleService := service.NewArticleService(articleRepo)
	ag := ArticleGroup{
		articleService: articleService,
		logger:         logger,
	}
	engine.POST("/articles", ag.SaveArticleHandler)
	engine.GET("/articles/:article_id", ag.FetchArticleByIDHandler)
	engine.GET("/articles", ag.FetchArticlesHandler)
	return engine, nil
}
