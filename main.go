package main

import (
	"log"
	"net/http"

	"github.com/baalamurgan/crud-golang/pkg/config"
	"github.com/baalamurgan/crud-golang/pkg/routes"
	"github.com/gorilla/mux"
)

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	config.Connect()
	route := mux.NewRouter()
	routes.RegisterWeatherRoutes(route)
	routes.PostRoutes(route)
	http.Handle("/", route)
	Port := "8001"
	log.Println("started server Running at http://localhost:" + Port) // this is just hang

	log.Fatal(http.ListenAndServe(":"+Port, setHeaders(route))) // this is reall prot
	//
}
