package main

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

type Config struct {
	Port int
	MavenDir string
}

func main() {

	// top level error handling
	defer ErrorHandler()

	// parse config
	var conf Config
	if _, err := toml.DecodeFile("moddex.conf", &conf); err != nil {
		panic(err)
	}

	// get gorilla router
	router := mux.NewRouter()

	// global middleware for logging and panic recovery
	n := negroni.New(negroni.NewLogger(), negroni.NewRecovery())

	MavenInit(conf, router.PathPrefix("/maven").Subrouter())
	router.PathPrefix("/rest/v0.1").Subrouter() // TODO: rest API

	// serve angular properly
	router.PathPrefix("/").Handler(negroni.New(
		negroni.NewStatic(http.Dir("./angular/")),
//		negroni.HandlerFunc(Index),
	))

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

func Index(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	http.ServeFile(rw, r, "angular/index.html")
}
