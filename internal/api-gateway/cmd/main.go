package main

import (
	"log"
	"net/http"

	grpc_client "imantask/internal/api-gateway/controller/grpc"
	handler "imantask/internal/api-gateway/controller/rest"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	auditClient, err := grpc_client.NewClient(5050)
	log.Println("sdf")
	if err != nil {
		log.Println(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	postHandler := handler.NewHandler(*auditClient)

	// Testing router
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	log.Println("ssfd")

	router.Route("/post", func(r chi.Router) {
		r.Get("/", postHandler.SavePosts)
	})
}
