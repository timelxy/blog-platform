package handler

import (
	"blog-platform/model"
	"blog-platform/service"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestHandlerImpl_RetrieveAllPosts(t *testing.T) {
	// echo and context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)

	mockService := service.NewMockPostServiceInterface(ctrl)

	mockService.EXPECT().RetrieveAllPosts(gomock.Any()).Return([]model.Post{}, nil).Times(1)
	mockService.EXPECT().RetrieveAllPosts(gomock.Any()).Return(nil, errors.New("fail!")).Times(1)

	type fields struct {
		Service service.PostServiceInterface
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				Service: mockService,
			},
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				Service: mockService,
			},
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				Service: tt.fields.Service,
			}
			if err := h.RetrieveAllPosts(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.RetrieveAllPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandlerImpl_CreatePost(t *testing.T) {
	jsonData := []byte(`{"title":"Test Post","content":"This is a test post content."}`)

	// echo and context
	e := echo.New()

	ctrl := gomock.NewController(t)

	mockService := service.NewMockPostServiceInterface(ctrl)

	mockService.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	mockService.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(errors.New("fail!")).Times(1)

	type fields struct {
		Service service.PostServiceInterface
	}
	type args struct {
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		res     model.Post
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				Service: mockService,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				Service: mockService,
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(jsonData))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			h := &HandlerImpl{
				Service: tt.fields.Service,
			}
			if err := h.CreatePost(ctx); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
