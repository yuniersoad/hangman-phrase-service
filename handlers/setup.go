package handlers

import (
	"encoding/json"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/yuniersoad/hangman-phrase-service/storage"
	"net/http"
)

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/phrases/random", GetRandomPhrase)
	r.Get("/phrases", GetPhrases)
	r.Post("/phrases", AddPhrase)
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

func GetPhrases(w http.ResponseWriter, r *http.Request) {
	phrases, err := storage.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(phrases)
}

func AddPhrase(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := storage.Add(text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
