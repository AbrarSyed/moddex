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
	// parse config
	var conf config
	if _, err := toml.DecodeFile("moddex.conf", &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(" port -> %d \n", conf.Port)
	fmt.Printf(" maven dir -> %s \n", conf.MavenDir)

	// get gorilla router
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
	n.Run(fmt.Sprint(":", conf.Port))
}
