package main

import (
	"imantask/common"
	"log"
	"os"

	postAdapter "imantask/internal/collector/adapters"
	postApp "imantask/internal/collector/app"

	"time"

	server "imantask/internal/collector/server"
)

func main() {
	grpcServer()
}

func grpcServer() {
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
	postProvider, err := postAdapter.NewClient(time.Second * 25)
	if err != nil {
		log.Println(err)
	}
	postService := postApp.NewPostService(postRepo, postProvider)
	postServer := server.NewPostServer(postService)

	srv := server.New(&postServer)

	log.Println("SERVER STARTED", time.Now())

	if err != nil {
		log.Println(err)
	}
	if err := srv.ListenAndServe(5050); err != nil {
		log.Println(err)
	}

	log.Println("listening on port: ", 5050)

	log.Println("sdf")

}
