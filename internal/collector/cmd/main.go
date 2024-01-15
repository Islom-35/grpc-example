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

	collectorRepo := postAdapter.NewCollectorRepository(db)
	postProvider, err := postAdapter.NewClient(time.Second * 10)
	if err != nil {
		log.Println(err)
	}
	collectorService := postApp.NewCollectorService(collectorRepo, postProvider)
	collectorServer := server.NewCollectorServer(collectorService)

	srv := server.New(&collectorServer)

	log.Println("SERVER STARTED", time.Now())

	if err != nil {
		log.Println(err)
	}
	port :=os.Getenv("COLLECTOR_PORT")
	if err := srv.ListenAndServe(port); err != nil {
		log.Println(err)
	}


	// log.Println("listening on port: ", 5050)

}
