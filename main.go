package main

import (
	"bufunfa/middlewares"
	"bufunfa/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Request received:", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	router.Group("/api/v1", func(r *routes.Router) {
		routes.RegisterAuthRoutes(r)

		r.Use(middlewares.AuthMiddleware)

		routes.RegisterUserRoutes(r)
		routes.RegisterProfileRoutes(r)
	})

	fmt.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
