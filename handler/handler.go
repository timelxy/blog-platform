package handler

import (
	"blog-platform/model"
	"blog-platform/resource"
	"blog-platform/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlerInterface interface {
	CreatePost(ctx echo.Context) (err error)
	RetrievePostByID(ctx echo.Context) (err error)
	RetrieveAllPosts(ctx echo.Context) (err error)
	HealthCheckHandler(ctx echo.Context) (err error)
}

type HandlerImpl struct {
	Service service.PostServiceInterface
}

func NewHandler() HandlerInterface {
	return &HandlerImpl{
		Service: service.NewPostService(),
	}
}

// RetrieveAllPosts godoc
// @Summary Retrieve all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} Post
// @Router /posts [get]
func (h *HandlerImpl) RetrieveAllPosts(ctx echo.Context) (err error) {
	posts, err := h.Service.RetrieveAllPosts(ctx)
	if err != nil {
		return
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
func (h *HandlerImpl) RetrievePostByID(ctx echo.Context) (err error) {
	id := ctx.Param("id")

	post, err := h.Service.RetrievePostByID(ctx, id)
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
func (h *HandlerImpl) CreatePost(ctx echo.Context) (err error) {
	// Parse query payload
	post := new(model.Post)
	if err = ctx.Bind(post); err != nil {
		return
	}
	createAt := time.Now()
	post.CreatedDate = createAt

	err = h.Service.CreatePost(ctx, post)
	if err != nil {
		return
	}

	return ctx.JSON(http.StatusOK, post)
}

// HealthCheckHandler godoc
// @Summary Check server health status
// @Description Check if the server is running and MongoDB connection is established
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "Everything is fine!"
// @Failure 500 {string} string "MongoDB connection failed!"
// @Router /healthcheck [get]
func (h *HandlerImpl) HealthCheckHandler(ctx echo.Context) (err error) {
	// Check mongodb connection
	err = resource.MongoClient.Ping(ctx.Request().Context(), nil)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "MongoDB connection failed!")
	}

	return ctx.String(http.StatusOK, "Everything is fine!")

}
