package main

import (
	"apica/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cache := handler.NewLRUCache(3)
	cache.Set(1, 1)
	cache.Set(2, 2)
	router := chi.NewRouter()
	router.Route("/lru", func(r chi.Router) {
		r.Get("/{key}", cache.Get)
	})
	fmt.Printf("server started at: %s", "3009")
	http.ListenAndServe(":3009", router)
}
