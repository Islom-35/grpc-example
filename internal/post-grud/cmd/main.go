package main

import (
	"imantask/common"
	postAdapter "imantask/internal/post-grud/adapters"
	"imantask/internal/post-grud/server"
	postApp "imantask/internal/post-grud/app"
	"log"
	"os"
	"time"
)

func main() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	postRepo := postAdapter.NewPostRepository(db)
	if err != nil {
		log.Println(err)
	}
	postService := postApp.NewPostService(postRepo)
	postServer := server.NewPostServer(postService)

	srv := server.New(&postServer)

	log.Println("SERVER STARTED", time.Now())

	if err != nil {
		log.Println(err)
	}
	if err := srv.ListenAndServe(5040); err != nil {
		log.Println(err)
	}

	log.Println("listening on port: ", 5040)
}
