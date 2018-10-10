package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kskitek/k8sFirstSteps/influx"
	"github.com/kskitek/k8sFirstSteps/value"
	log "github.com/sirupsen/logrus"
)

var vg value.Generator
var vs value.Saver

func main() {
	port := flag.Int("port", 8080, "http server port")
	influxAddr := flag.String("influxAddr", "http://influx:8086", "InfluxDB address")
	influxUser := flag.String("influxUser", "", "InfluxDB username")
	influxPwd := flag.String("influxPwd", "", "InfluxDB password")
	influxDB := flag.String("influxDB", "mes", "InfluxDB database")
	flag.Parse()

	serverAddr, router := setupHTTPServer(port)

	setupGenerator(*influxAddr, *influxUser, *influxPwd, *influxDB)

	log.Infof("Starting Server at port %d.", *port)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}

func setupGenerator(addr, user, pwd, db string) {
	var err error
	vs, err = influx.New(addr, user, pwd, db, 1)
	if err != nil {
		log.Fatal(err)
	}

	vg = value.New(10)
	ch := vg.Start(1)
	go func() {
		for {
			v := <-ch
			vs.Save(v)
		}
	}()
	log.Info("Generator set up")
}

func setupHTTPServer(port *int) (string, *httprouter.Router) {
	router := httprouter.New()
	serverAddr := fmt.Sprintf(":%d", *port)

	router.GET("/set/:value", loggingHandler(set))
	router.GET("/healthz", healthz)
	router.NotFound = notFound
	log.Info("HTTP Server set up")

	return serverAddr, router
}

func set(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	value, err := intFromParams(params)
	if err != nil {
		log.Warnf("wrong value %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		log.Infof("new value %d", value)
		vg.Value(value)
		w.WriteHeader(http.StatusOK)
	}
}

func intFromParams(p httprouter.Params) (int, error) {
	s := p.ByName("value")
	return strconv.Atoi(s)
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
