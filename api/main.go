package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"os"
	"wallet-go/api/handle"
	"wallet-go/api/middleware"
	"wallet-go/src/usecase/healthcheck"

	"github.com/gorilla/mux"
)

var port string

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		port = "8585"
	}

}

func main() {
	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
	)

	handle.MakeHealthHandlers(r, *n, healthcheck.Service{})
	listen := fmt.Sprintf("0.0.0.0:%v", port)
	fmt.Println("Listen service in port", listen)
	log.Fatalln(http.ListenAndServe(listen, r))
}
