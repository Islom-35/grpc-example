package main

import (
	"imantask/common"
	postAdapter "imantask/internal/post/adapters"
	postApp "imantask/internal/post/app"
	"imantask/internal/post/server"
	"log"
	"os"
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

	if err != nil {
		log.Println(err)
	}
	if err := srv.ListenAndServe(os.Getenv("POST_PORT")); err != nil {
		log.Println(err)
	}

	log.Println("listening on port: ", 5040)
}
