package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ensena/auth-token"
)

func HttpServer() {

	http.Handle("/api/v1/info", auth.MiddleWare(http.HandlerFunc(getInfo)))
	http.HandleFunc("/api/v1/news", getInfo)
	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
