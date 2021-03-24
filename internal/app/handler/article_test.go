package handler

import (
	"assignment/internal/app/model"
	"context"

	db "assignment/db"
	repo "assignment/internal/app/repo"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

var articleHandler http.Handler
var conn *mongo.Database
var ctx context.Context

func init() {
	logFile, err := os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger := log.New(logFile, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	mongoURI := "mongodb://localhost:27017/"
	dbName := "ArticleTestDB"
	conn, err = db.Connect(mongoURI, dbName)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	var ctx context.Context
	defer conn.Drop(ctx)
	articleHandler, err = New(conn, logger)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}

func TestArticleGroup_SaveArticleHandler(t *testing.T) {
	defer repo.TruncateCollection(conn)

	tests := []struct {
		name    string
		article model.Article
		status  int
	}{
		{
			name: "createArticleSuccess",
			article: model.Article{
				Title:   "Hello World",
				Content: "This is for testing",
				Author:  "menka",
			},
			status: 201,
		},
		{
			name: "createArticleFail",
			article: model.Article{
				Title:  "Hello Reader",
				Author: "menka",
			},
			status: 422,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			articleByte, err := json.Marshal(tt.article)
			if err != nil {
				t.Fatalf("test fail: expected status %d got marshing error", tt.status)
			}
			req := httptest.NewRequest("POST", "/articles", bytes.NewReader(articleByte))
			req.Header.Set("Content-Type", "application/json")
			articleHandler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
			t.Logf("test pass: expected status %d, got %d", tt.status, rr.Code)
		})
	}
}

func TestArticleGroup_FetchArticleByIDHandler(t *testing.T) {
	defer repo.TruncateCollection(conn)
	articleList, err := repo.SeedDummyArticle(conn)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name     string
		articlID string
		status   int
	}{
		{
			name:     "getArticleByIDSuccess",
			articlID: articleList[0].ID,
			status:   200,
		},
		{
			name:     "getArticleByIDError",
			articlID: "1234",
			status:   500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/articles/"+tt.articlID, nil)
			articleHandler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
			t.Logf("test pass: expected status %d, got %d", tt.status, rr.Code)
		})
	}
}

func TestArticleGroup_FetchArticlesHandler(t *testing.T) {
	defer repo.TruncateCollection(conn)
	_, err := repo.SeedDummyArticle(conn)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "getArticleSuccess",
			status: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/articles", nil)
			articleHandler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
			t.Logf("test pass: expected status %d, got %d", tt.status, rr.Code)
		})
	}
}
