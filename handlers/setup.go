package handlers

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"net/http"
)

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/phrases/random", GetRandomPhrase)
	http.Handle("/", r)
}

func GetRandomPhrase(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("al que madruga dios lo ayuda"))
}
