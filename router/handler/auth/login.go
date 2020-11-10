package auth

import (
	"fmt"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("SIGN UP")))
}
