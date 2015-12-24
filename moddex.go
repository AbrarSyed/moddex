package main

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

type config struct {
	Port int
	MavenDir string
}

func main() {

	// top level error handling
	defer ErrorHandler()

	// parse config
	var conf config
	if _, err := toml.DecodeFile("moddex.conf", &conf); err != nil {
		panic(err)
	}

	// get gorilla router
	router := mux.NewRouter()

	// global middleware for logging and panic recovery
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery())

	router.PathPrefix("maven").Subrouter() // TODO: maven router
	router.PathPrefix("rest/v0.1").Subrouter() // TODO: rest API

	// redirect root to web
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		http.Redirect(resp, req, "/web/", 301)
	})

	// angular web
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web", http.FileServer(http.Dir("./angular/"))))

	// launch!
	n.UseHandler(router)
	n.Run(fmt.Sprint(":", conf.Port))
}

func ErrorHandler() {
	if err := recover(); err != nil {
		fmt.Println("Moddex died: ", err)
		os.Exit(1)
	}
}
