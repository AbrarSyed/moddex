package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

func MavenInit(conf Config, router *mux.Router) {

	// handle GET requests
	router.Methods("GET", "HEAD").Handler(http.StripPrefix("/maven/", negroni.New(
		negroni.NewStatic(http.Dir(conf.MavenDir)),
//		negroni.HandlerFunc(indexHandler),
	)))
}

//func handleFile(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//	// TODO: show file index
//	http.ServeFile(rw, r, "angular/index.html")
//}
//
//func indexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//	// TODO: show file index
//	http.ServeFile(rw, r, "angular/index.html")
//}