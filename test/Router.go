package test

import (
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func Router(path string, handler Handler) {
	http.HandleFunc(path, handler)
	fmt.Println("server is listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
