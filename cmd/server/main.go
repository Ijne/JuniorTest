package main

import (
	"log"
	"net/http"

	"github.com/Ijne/JuniorTest/config"
	"github.com/Ijne/JuniorTest/internal/handlers"
	"github.com/Ijne/JuniorTest/internal/storage/postgres"
	"github.com/go-chi/chi"
)

func main() {
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	s, err := postgres.NewStorage(cfg)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	r := chi.NewRouter()

	r.Post("/createsub", handlers.CreateSubHandler(s))
	r.Handle("/getsub", handlers.GetSubHandler(s))
	r.Handle("/updatesub", handlers.UpdateSubHandler(s))
	r.Handle("/deletesub", handlers.DeleteSubHandler(s))
	r.Handle("/list", handlers.ListSubHandler(s))
	r.Handle("/amount", handlers.AmountSubHandler(s))

	log.Println("Start serving...")
	if err := http.ListenAndServe("localhost:8087", r); err != nil {
		panic(err)
	}
}
