package handlers

import (
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !"))
}
