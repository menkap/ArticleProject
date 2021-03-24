package model

//Article
type Article struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Title   string `json:"title" bson:"title" validate:"required"`
	Content string `json:"content" bson:"content" validate:"required"`
	Author  string `json:"author" bson:"author" validate:"required"`
}
