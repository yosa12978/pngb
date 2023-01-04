package repositories

import (
	"context"

	"github.com/yosa12978/pngb/internal/pkg/helpers"
	"github.com/yosa12978/pngb/internal/pkg/models"
	"github.com/yosa12978/pngb/internal/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository interface {
	Find() []models.Post
	FindByID(id string) (models.Post, error)
	Create(p models.Post) error
	Update(p models.Post) error
	Delete(id string) error
}

type postRepositoryMongo struct {
}

func NewPostMongoRepo() PostRepository {
	return new(postRepositoryMongo)
}

func (repo *postRepositoryMongo) Find() []models.Post {
	db := mongodb.Connect()
	var posts []models.Post
	opts := options.Find().SetSort(bson.M{"_id": -1})
	cursor, err := db.Collection("posts").Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return posts
	}
	cursor.Decode(&posts)
	return posts
}

func (repo *postRepositoryMongo) FindByID(id string) (models.Post, error) {
	db := mongodb.Connect()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Post{}, helpers.ErrNotFound
	}
	var post models.Post
	err = db.Collection("posts").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return post, helpers.ErrNotFound
	}
	return post, nil
}

func (repo *postRepositoryMongo) Create(p models.Post) error {
	db := mongodb.Connect()
	p.Id = primitive.NewObjectID()
	_, err := db.Collection("posts").InsertOne(context.TODO(), p)
	if err != nil {
		return helpers.ErrBadRequest
	}
	return nil
}

func (repo *postRepositoryMongo) Update(p models.Post) error {
	db := mongodb.Connect()
	_, err := db.Collection("posts").ReplaceOne(context.TODO(), p, bson.M{"_id": p.Id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helpers.ErrNotFound
		}
		return helpers.ErrBadRequest
	}
	return nil
}

func (repo *postRepositoryMongo) Delete(id string) error {
	db := mongodb.Connect()
	_, err := db.Collection("posts").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return helpers.ErrBadRequest
	}
	return nil
}
