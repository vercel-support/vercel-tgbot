package api

import (
	"fmt"
	"net/http"
)

func ReplyHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World! Welcome: %s", r.URL.Path[1:])
}
