package service

//go:generate mockgen -source=post.go -destination=./mock_post.go -package=service

import (
	"blog-platform/model"
	"blog-platform/resource"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostServiceInterface interface {
	CreatePost(ctx echo.Context, post *model.Post) (err error)
	RetrieveAllPosts(ctx echo.Context) (posts []model.Post, err error)
	RetrievePostByID(ctx echo.Context, id string) (post model.Post, err error)
}

type PostServiceImpl struct {
	Collection *mongo.Collection
}

func NewPostService() PostServiceInterface {
	collection := resource.MongoClient.Database("blog").Collection("posts")
	return &PostServiceImpl{
		Collection: collection,
	}
}

// CreatePost
func (ps *PostServiceImpl) CreatePost(ctx echo.Context, post *model.Post) (err error) {
	insertResult, err := ps.Collection.InsertOne(context.Background(), post)
	if err != nil {
		return
	}

	id, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("insertResult of mongo has no InsertedID")
	}

	// inserted document id
	post.ID = &id

	return
}

// RetrieveAllPosts
func (ps *PostServiceImpl) RetrieveAllPosts(ctx echo.Context) (posts []model.Post, err error) {
	cursor, err := ps.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())

	// Traverse all posts
	for cursor.Next(context.Background()) {
		var post model.Post
		if err = cursor.Decode(&post); err != nil {
			return
		}
		posts = append(posts, post)
	}

	return
}

// RetrievePostByID
func (ps *PostServiceImpl) RetrievePostByID(ctx echo.Context, id string) (post model.Post, err error) {
	// id string to objectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	// Condition
	filter := bson.M{"_id": objectID}

	// Find it
	err = ps.Collection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return
	}

	return
}
