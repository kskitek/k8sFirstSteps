package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()

	serverAddr, router := setupHTTPServer(port)

	log.Infof("Starting Server at port %d. Ip Addr ?.", *port)
	log.Fatal(http.ListenAndServe(serverAddr, router))
	os.Exit(1)
}

func setupHTTPServer(port *int) (string, *httprouter.Router) {
	router := httprouter.New()
	serverAddr := fmt.Sprintf(":%d", *port)

	router.GET("/set", loggingHandler(set))
	router.GET("/healthz", healthz)
	router.NotFound = notFound

	return serverAddr, router
}

func set(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func healthz(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func notFound(w http.ResponseWriter, req *http.Request) {
	log.Warnf("404 %s %s", req.Method, req.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

func loggingHandler(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Infof("%s %s", req.Method, req.RequestURI)
		handler(w, req, params)
	}
}
