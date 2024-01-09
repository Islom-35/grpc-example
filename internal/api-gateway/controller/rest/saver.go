package rest

import (
	"context"
	"log"
	"net/http"
	"time"
	grpc_client "imantask/internal/api-gateway/controller/grpc"
)



type Handler struct {
	postsService grpc_client.Client
}

func NewHandler(posts grpc_client.Client) *Handler {
	return &Handler{
		postsService: posts,
	}
}

func (h *Handler) SavePosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.postsService.Save(ctx)

	if err != nil {
		log.Println("CreateOrder() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
