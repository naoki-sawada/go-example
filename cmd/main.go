package main

import (
	"go-example/internal/interfaces/api/server"
	"go-example/internal/platform/database"
	"go-example/internal/registry"
)

func main() {
	db, err := database.Connect(database.ReadConfig())
	if err != nil {
		panic(err)
	}
	r := registry.NewRegistry(db)
	s := server.Create(r)
	s.Run(":8080")
}
