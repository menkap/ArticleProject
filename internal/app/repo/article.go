package repo

import (
	"assignment/internal/app/model"
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CollectionArticles -
const CollectionArticles = "articles"

// ArticleRepoI - interface of product repository
type ArticleRepoI interface {
	Create(ctx context.Context, articleObj model.Article) (string, error)
	FindByID(ctx context.Context, ID string, articleView *model.Article) error
	FetchAll(ctx context.Context, articleViews *[]model.Article) error
}

type article struct {
	conn *mongo.Database
}

//NewArticleRepo - return instance of article repository
func NewArticleRepo(connObj *mongo.Database) (ArticleRepoI, error) {
	if connObj == nil {
		return nil, errors.New("please provide connection instance")
	}
	return article{
		connObj,
	}, nil
}

// Create - add articles into database
func (repo article) Create(ctx context.Context, articleObj model.Article) (string, error) {
	CollectionArticles := repo.conn.Collection(CollectionArticles)

	instRes, err := CollectionArticles.InsertOne(ctx, articleObj)
	if err != nil {
		return "", err
	}
	ID, isObjectID := instRes.InsertedID.(primitive.ObjectID)
	if !isObjectID {
		return "", errors.New("fail to fetch ID")
	}

	return ID.Hex(), nil
}

//FindByID - fetch article by ID from database
func (repo article) FindByID(ctx context.Context, ID string, articleView *model.Article) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return errors.New("invalid ID")
	}
	selector := bson.D{{"_id", objID}}
	CollectionArticles := repo.conn.Collection(CollectionArticles)
	err = CollectionArticles.FindOne(ctx, selector).Decode(&articleView)
	return err
}

// FetchAll - fetch articles from database
func (repo article) FetchAll(ctx context.Context, articleViews *[]model.Article) error {
	selector := bson.D{{}}
	CollectionArticles := repo.conn.Collection(CollectionArticles)
	cur, err := CollectionArticles.Find(ctx, selector)
	if err != nil {
		return err
	}
	if err := cur.All(ctx, articleViews); err != nil {
		return err
	}
	return nil
}

// SeedDummyArticle - adds dummy article object in db for testing
func SeedDummyArticle(conn *mongo.Database) ([]model.Article, error) {
	articleList := []model.Article{
		{
			Title:   "New dummy article",
			Content: "Dummy articles to be published",
			Author:  "Menka",
		},
		{
			Title:   "New dummy article2",
			Content: "Dummy articles to be published",
			Author:  "Menka",
		},
	}
	CollectionArticles := conn.Collection(CollectionArticles)
	for idx, articleObj := range articleList {

		instRes, err := CollectionArticles.InsertOne(context.Background(), articleObj)
		if err != nil {
			return nil, err
		}
		ID, isObjectID := instRes.InsertedID.(primitive.ObjectID)
		if !isObjectID {
			return nil, errors.New("fail to fetch ID")
		}

		articleList[idx].ID = ID.Hex()

	}
	return articleList, nil
}

// TruncateCollection - truncates article collection
func TruncateCollection(db *mongo.Database) {
	db.Collection(CollectionArticles).DeleteMany(context.Background(), bson.D{})
}

// MockArticleRepo - Mock type for article repo
type MockArticleRepo struct {
	mock.Mock
}

//Create - mock function of article repo Create
func (m *MockArticleRepo) Create(ctx context.Context, articleObj model.Article) (string, error) {
	ret := m.Called(ctx, articleObj)
	return ret.String(0), ret.Error(1)
}

//FindByID - mock function of article repo FindByID
func (m *MockArticleRepo) FindByID(ctx context.Context, ID string, articleView *model.Article) error {
	ret := m.Called(ctx, articleView)
	return ret.Error(0)
}

// FetchAll - mock function of article repo FetchAll
func (m *MockArticleRepo) FetchAll(ctx context.Context, articleViews *[]model.Article) error {
	ret := m.Called(ctx, articleViews)
	return ret.Error(0)
}
