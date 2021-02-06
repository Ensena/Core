package router

import (
	"net/http"

	"github.com/Ensena/core"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Write(core.GetInfo(1))
}
