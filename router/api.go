package router

import (
	"go-docker/helper"
	"go-docker/router/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func API(r *mux.Router) *mux.Router {
	// /api/v1/
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	v1.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
	v1.Use(middleware.ApiAccessMiddleware)
	// /api/v1/ping/
	registerRoute(v1, ApiRoute["v1"])
	// /api/v1/auth/
	v1Auth := v1.PathPrefix("/auth").Subrouter()
	registerRoute(v1Auth, ApiRoute["auth"])
	return r
}
func registerRoute(r *mux.Router, routes []Route) {
	for _, route := range routes {
		r.Methods(route.Method).Path(route.Path).HandlerFunc(route.Handler).Name(route.Name)
	}
}
func handlePing(w http.ResponseWriter, r *http.Request) {
	helper.ReturnResponseAsJSON(w, nil, "PONG", 404)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	helper.ReturnResponseAsJSON(w, nil, "Not-Found", 404)
	return
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	helper.ReturnResponseAsJSON(w, nil, "Method Not Allowd", 405)
	return
}

func FetchAllRoute() {
	routes := []interface{}{}
	for _, v := range ApiRoute {
		routes = append(routes, v)
	}
	logrus.Infoln(routes)
}
