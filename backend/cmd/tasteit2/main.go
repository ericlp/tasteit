package main

import (
	"github.com/ericlp/tasteit2/backend/internal/api"
	"github.com/ericlp/tasteit2/backend/internal/db"
	"log"
)

func main() {
	log.Println("==== Starting tasteit2 golang backend =====")

	db.Init()
	defer db.Close()
	api.Init()
	api.Start()
}
