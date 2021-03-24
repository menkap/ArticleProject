package service

import (
	"assignment/internal/app/model"
	"assignment/internal/app/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSaveArticleService(t *testing.T) {
	mockArticleRepo := new(repo.MockArticleRepo)
	tests := []struct {
		name     string
		article  model.Article
		initFunc func()
		wantErr  bool
	}{
		{
			name: "createArticleOne",
			article: model.Article{
				Title:   "Hello World",
				Content: "This is for testing",
				Author:  "menka",
			},
			initFunc: func() {
				mockArticleRepo.On("Create", mock.Anything, mock.Anything).Return("123", nil)
			},
			wantErr: false,
		},
	}

	articleService := NewArticleService(mockArticleRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.initFunc != nil {
				tt.initFunc()
			}
			_, err := articleService.SaveArticleService(context.Background(), tt.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.SaveArticleService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("test pass")
		})
	}
}

func TestFetchArticleByIDService(t *testing.T) {
	tests := []struct {
		name         string
		articleID    string
		articleViews *model.Article
		wantErr      bool
		mockReturns  struct {
			err        error
			articlePtr *model.Article
		}
	}{
		{
			name:         "getArticleByIDSuccess",
			articleID:    "1",
			articleViews: &model.Article{},
			mockReturns: struct {
				err        error
				articlePtr *model.Article
			}{
				err: nil,
				articlePtr: &model.Article{
					ID: "1",
				},
			},
			wantErr: false,
		},
		{
			name:         "getArticleByIDFail",
			articleID:    "5",
			articleViews: &model.Article{},
			mockReturns: struct {
				err        error
				articlePtr *model.Article
			}{
				err:        errors.New("not found"),
				articlePtr: nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockArticleRepo := new(repo.MockArticleRepo)
			mockArticleRepo.On("FindByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockReturns.err, tt.mockReturns.articlePtr)
			articleService := NewArticleService(mockArticleRepo)
			err := articleService.FetchArticleByIDService(context.Background(), tt.articleID, tt.articleViews)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.SaveArticleService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("test pass")
		})
	}
}

func TestFetchArticlesService(t *testing.T) {
	tests := []struct {
		name         string
		articleViews *[]model.Article
		mockErr      error
		wantErr      bool
	}{
		{
			name:         "getArticlesSuccess",
			articleViews: &[]model.Article{},
			mockErr:      nil,
			wantErr:      false,
		},
		{
			name:         "getArticlesFail",
			articleViews: &[]model.Article{},
			mockErr:      errors.New("not found"),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockArticleRepo := new(repo.MockArticleRepo)
			mockArticleRepo.On("FetchAll", mock.Anything, mock.Anything).Return(tt.mockErr)
			articleService := NewArticleService(mockArticleRepo)
			err := articleService.FetchArticlesService(context.Background(), tt.articleViews)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.SaveArticleService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("test pass")
		})
	}
}
