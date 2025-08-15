package routes

import (
	"bufunfa/middlewares"
	"fmt"
	"net/http"
)

func RegisterProfileRoutes(r *Router) {
	r.GET("/profile", func(w http.ResponseWriter, r *http.Request) {
		userId := middlewares.GetUserID(r)
		fmt.Fprintln(w, "Perfil do usuário", userId)
	})

	r.PUT("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Atualizar perfil")
	})
}
