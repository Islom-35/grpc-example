package main

import (
	"log"
	"net/http"
	"os"

	grpc_client "imantask/internal/api-gateway/controller/grpc"
	handler "imantask/internal/api-gateway/controller/rest"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	auditClient, err := grpc_client.NewClient(os.Getenv("COLLECTOR_PORT"), os.Getenv("POST_PORT"))
	if err != nil {
		log.Println(err, "port")
	}
	defer auditClient.CloseConnection()

	postHandler := handler.NewHandler(*auditClient)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Testing router
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	

	router.Route("/post", func(r chi.Router) {
		r.Post("/", postHandler.CollectPosts)
		r.Get("/{id:[0-9]+}", postHandler.GetPostByID)
		r.Get("/", postHandler.GetPage)
		r.Delete("/{id:[0-9]+}", postHandler.DeletePostByID)
	})

	gPort :=(":5060")
	err = http.ListenAndServe(gPort, router)

	log.Println("listening on port: ", 5060)
	log.Println(">>>")
	if err != nil {
		
		log.Fatal(err)
	}

}