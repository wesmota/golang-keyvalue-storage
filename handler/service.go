package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wesmota/keyvalue-storage/core"
)

func keyValueGetHandler(r *mux.Router, storage *core.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		value, err := storage.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(value))
	}
}

func KeyValuePutHandler(r *mux.Router, storage *core.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		value, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = storage.Put(key, string(value))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		log.Printf("PUT key=%s value=%s\n", key, string(value))
	}
}

func KeyValueDeleteHandler(r *mux.Router, storage *core.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		err := storage.Delete(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

func MakeServiceHandlers(r *mux.Router, storage *core.Storage) {

	r.HandleFunc("/v1/{key}", keyValueGetHandler(r, storage)).Methods("GET")
	r.HandleFunc("/v1/{key}", KeyValuePutHandler(r, storage)).Methods("PUT")
	r.HandleFunc("/v1/{key}", KeyValueDeleteHandler(r, storage)).Methods("DELETE")
	r.HandleFunc("/v1", notAllowedHandler)
	r.HandleFunc("/v1/{key}", notAllowedHandler)

}
