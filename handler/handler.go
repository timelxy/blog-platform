package handler

import (
	"blog-platform/resource"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title"`
	Content     string              `json:"content,omitempty" bson:"content"`
	CreatedDate time.Time           `json:"created_date,omitempty" bson:"cteated_date"`
}

// RetrieveAllPosts godoc
// @Summary Retrieve all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} Post
// @Router /posts [get]
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

// RetrievePostByID godoc
// @Summary Retrieve a post by ID
// @Description Get a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} Post
// @Router /posts/{id} [get]
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

// CreateNewPost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param input body Post true "Post object that needs to be added"
// @Success 200 {object} Post
// @Router /posts [post]
func CreateNewPost(ctx echo.Context) (err error) {
	// Parse query payload
	payload := new(Post)
	if err = ctx.Bind(payload); err != nil {
		return
	}

	createAt := time.Now()

	post := Post{
		Title:       payload.Title,
		Content:     payload.Content,
		CreatedDate: createAt,
	}

	// Insert
	collection := resource.MongoClient.Database("blog").Collection("posts")
	insertResult, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return
	}

	id, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("insertResult of mongo has no InsertedID")
	}

	output := Post{
		ID:          &id,
		CreatedDate: createAt,
	}

	return ctx.JSON(http.StatusOK, output)
}
