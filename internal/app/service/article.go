package service

import (
	"assignment/internal/app/model"
	"assignment/internal/app/repo"
	"context"

	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleService -
type ArticleService struct {
	articleRepo repo.ArticleRepoI
}

//NewArticleService - creates instance of article service
func NewArticleService(repo repo.ArticleRepoI) *ArticleService {
	service := ArticleService{
		repo,
	}
	return &service
}

//SaveArticleService - add articles into database
func (s *ArticleService) SaveArticleService(ctx context.Context, articleObj model.Article) (string, error) {
	ID, err := s.articleRepo.Create(ctx, articleObj)
	if err != nil {
		return "", errors.New("failed to save user details: " + err.Error())
	}
	return ID, nil
}

// FetchArticleByIDService - fetch article by ID from database
func (s *ArticleService) FetchArticleByIDService(ctx context.Context, ID string, articleView *model.Article) error {
	err := s.articleRepo.FindByID(ctx, ID, articleView)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("not found")
		}
		return err
	}
	return nil
}

// FetchArticlesService - fetch articles from database
func (s *ArticleService) FetchArticlesService(ctx context.Context, articleViews *[]model.Article) error {
	err := s.articleRepo.FetchAll(ctx, articleViews)
	if err != nil {
		return err
	}
	return nil
}
