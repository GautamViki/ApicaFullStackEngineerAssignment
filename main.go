package main

import (
	"apica/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	cache := handler.NewLRUCache(3)
	router.Route("/lru", func(r chi.Router) {
		r.Get("/{key}", cache.Get)
		r.Post("/", cache.Set)
		r.Delete("/{key}", cache.Delete)
	})
	fmt.Println("server started at: 3009")
	http.ListenAndServe(":3009", router)
}
