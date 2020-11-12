package router

import (
	"fmt"
	"go-docker/helper"
	"go-docker/router/handler/auth"
	"go-docker/router/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func API(r *mux.Router) *mux.Router {
	// /api/v1/
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	v1.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
	v1.Use(middleware.ApiAccessMiddleware)
	// /api/v1/ping/
	v1.Methods("GET").Path("/ping").HandlerFunc(handlePing)

	// /api/v1/auth/
	v1Auth := v1.PathPrefix("/auth").Subrouter()
	v1Auth.Methods("POST").Path("/sign-up").HandlerFunc(auth.SignUpHandler).Name("v1.auth.sign-up")
	v1Auth.Methods("POST").Path("/sign-in").HandlerFunc(auth.SignInHandler).Name("v1.auth.sign-in")
	return r
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("PONG")))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	helper.ReturnResponseAsJSON(w, nil, "Not-Found", 404)
	return
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	helper.ReturnResponseAsJSON(w, nil, "Method Not Allowd", 405)
	return
}
