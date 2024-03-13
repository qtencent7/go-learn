package build_a_server

import (
	"go-learn/build_a_server/routes"
	"net/http"
)

var R *http.ServeMux = http.NewServeMux()

func Router() {
	R.HandleFunc("/", routes.GeneralHandler)
	R.HandleFunc("/uploadfile", routes.UploadFile)
}
