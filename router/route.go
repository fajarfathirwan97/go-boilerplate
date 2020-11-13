package router

import (
	"go-docker/router/handler/auth"
	"net/http"
)

type Route struct {
	Handler func(http.ResponseWriter, *http.Request)
	Path    string
	Method  string
	Name    string
}

var ApiRoute = map[string][]Route{
	"v1": {{Handler: handlePing, Path: "ping", Method: "GET"}},
	"auth": {
		{
			Handler: auth.SignUpHandler,
			Method:  "POST",
			Name:    "v1.auth.sign-up",
			Path:    "/sign-up",
		},
		{
			Handler: auth.SignInHandler,
			Method:  "POST",
			Name:    "v1.auth.sign-in",
			Path:    "/sign-in",
		},
	},
}
