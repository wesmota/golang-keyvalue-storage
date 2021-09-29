package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wesmota/keyvalue-storage/core"
	"github.com/wesmota/keyvalue-storage/handler"
)

func main() {
	r := mux.NewRouter()

	storage := core.NewStorage()
	handler.MakeServiceHandlers(r, storage)

	log.Fatal(http.ListenAndServe(":8080", r))
}
