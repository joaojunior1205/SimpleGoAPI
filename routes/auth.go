package routes

import (
	"bufunfa/controllers"
	"fmt"
	"net/http"
)

func RegisterAuthRoutes(r *Router) {
	r.POST("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginHandler(w, r)
	})

	r.POST("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Registro de usuário")
	})
}
