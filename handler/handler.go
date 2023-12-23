package handler

import (
	"context"
	"net/http"
	"ocr/resource"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedDate time.Time `json:"created_date"`
}

// Retrieve all posts
func RetrieveAllPosts(ctx echo.Context) (err error) {
	collection := resource.MongoClient.Database("blog").Collection("posts")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())

	// Traverse all posts
	var posts []Post
	for cursor.Next(context.Background()) {
		var post Post
		if err = cursor.Decode(&post); err != nil {
			return
		}
		posts = append(posts, post)
	}

	return ctx.JSON(http.StatusOK, posts)

}

// Retrieves a blog post with given ID
func RetrievePostByID(ctx echo.Context) (err error) {
	collection := resource.MongoClient.Database("blog").Collection("posts")
	id := ctx.Param("id")

	// id string to objectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	// Condition
	filter := bson.M{"_id": objectID}

	// Find it
	var post Post
	err = collection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return
	}

	return ctx.JSON(http.StatusOK, post)
}

// Creates a new blog post
func CreateNewPost(ctx echo.Context) (err error) {
	// Parse query payload
	payload := new(Post)
	if err = ctx.Bind(payload); err != nil {
		return
	}

	post := Post{
		Title:       payload.Title,
		Content:     payload.Content,
		CreatedDate: time.Now(),
	}

	// Insert
	collection := resource.MongoClient.Database("blog").Collection("posts")
	insertResult, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return
	}

	return ctx.JSON(http.StatusOK, insertResult)
}
