package routes

import (
	"net/http"
)

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
