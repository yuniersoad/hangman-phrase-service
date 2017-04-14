package handlers

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/yuniersoad/hangman-phrase-service/storage"
	"net/http"
)

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/phrases/random", GetRandomPhrase)
	http.Handle("/", r)
}

func GetRandomPhrase(w http.ResponseWriter, r *http.Request) {
	phrase, err := storage.GetRandom()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(phrase))
}
