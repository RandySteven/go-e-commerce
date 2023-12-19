package main

import (
	"log"

	"github.com/RandySteven/go-e-commerce.git/db/postgres"
	"github.com/RandySteven/go-e-commerce.git/pkg/configs"
	"github.com/gorilla/mux"
)

func main() {
	configPath, err := configs.ParseFlags()
	if err != nil {
		log.Fatal(err)
		return
	}
	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	repo, err := postgres.NewRepositories(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(repo)

	r := mux.NewRouter()
	config.Run(r)
}
