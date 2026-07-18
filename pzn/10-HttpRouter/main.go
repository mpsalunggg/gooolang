package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Helloo gais")
	})

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: router,
	}

	server.ListenAndServe()

}
