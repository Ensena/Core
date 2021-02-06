package router

import (
	"fmt"
	"net/http"

	"github.com/Ensena/core"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	core.GetInfo(2)
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}
