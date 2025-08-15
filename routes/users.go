package routes

import (
	"fmt"
	"net/http"
)

func RegisterUserRoutes(r *Router) {
	r.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Lista de usuários")
	})

	r.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Criar usuário")
	})

	r.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Detalhes do usuário")
	})
}
