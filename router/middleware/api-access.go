package middleware

import (
	"go-docker/config"
	"go-docker/helper"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ApiAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b64, err := helper.Base64DecodeStripped(r.Header.Get("Authorization"))
		if err != nil {
			helper.ReturnResponseAsJSON(w, nil, "Unauthorized", 403)
			return
		}
		b64Sep := strings.Split(b64, ":")
		if b64Sep[0] != config.GetEnv().AccessKey || b64Sep[1] != config.GetEnv().AccessSecret {
			helper.ReturnResponseAsJSON(w, nil, "Unauthorized", 403)
			return
		}
		log.Debugln(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
