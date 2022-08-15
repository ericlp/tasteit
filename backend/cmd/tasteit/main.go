package main

import (
	"github.com/ericlp/tasteit/backend/internal/api"
	"github.com/ericlp/tasteit/backend/internal/db"
	"log"
)

func main() {
	log.Println("==== Starting tasteit golang backend =====")

	db.Init()
	defer db.Close()
	api.Init()
	api.Start()
}
