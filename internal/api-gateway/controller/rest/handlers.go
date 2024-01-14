package rest

import (
	"context"
	"encoding/json"
	"errors"
	grpc_client "imantask/internal/api-gateway/controller/grpc"
	"imantask/internal/genproto/ppb"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type Page struct {
	PageNumber uint32 `json:"pagenumber"`
	PageSize uint32 `json:"pagesize"`
}

type Handler struct {
	postClient grpc_client.Client
}

func NewHandler(client grpc_client.Client) *Handler {
	return &Handler{
		postClient: client,
	}
}

func (h *Handler) CollectPosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.postClient.CollectPosts(ctx)
	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ID, err := getIdFromRequest(r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := ppb.ID{
		ID: ID,
	}

	post, err := h.postClient.GetPostByID(ctx, &id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(post)

	response, err := json.Marshal(post)
	if err != nil {
		log.Println("<<<<")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)

}

func (h *Handler) GetPage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		cancel()
		return
	}

	var page Page

	if err = json.Unmarshal(reqBytes, &page); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		cancel()
		return
	}
	pagenumber:=page.PageNumber
	pagesize :=page.PageSize

	pbPage:=ppb.Page{
		PageNumber: pagenumber,
		PageSize: pagesize,
	}


	pages, err := h.postClient.GetPage(ctx, &pbPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	response, err := json.Marshal(pages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		cancel()
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := getIdFromRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// cancel()
		return
	}
	log.Println(ID)
	id := ppb.ID{
		ID: ID,
	}


	err = h.postClient.DeletePostByID(ctx, &id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// cancel()
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getIdFromRequest(r *http.Request) (uint32, error) {
	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if id64 == 0 {
		return 0, errors.New("id can't be 0")
	}
	id := uint32(id64)

	return id, nil
}
