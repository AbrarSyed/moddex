package main

import (
//	"log"
	"net/http"
	//"encoding/json"
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

type thing struct {
	data string
}

func main() {
	router := mux.NewRouter()

	// global stuff
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery())

	router.PathPrefix("maven").Subrouter() // maven router
	router.PathPrefix("rest/v0.1").Subrouter() // rest API

	// redirect root to web
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		http.Redirect(resp, req, "/web/", 301)
	})

	// angular web
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web", http.FileServer(http.Dir("./angular/"))))

	// launch!
	n.UseHandler(router)
	n.Run(":8080")
}
