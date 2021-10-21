package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

func NewServer(generator *Generator) Server {
	var config serverConfig
	err := envconfig.Process("SERVER", &config)
	if err != nil {
		panic(err)
	}

	return Server{
		generator: generator,
		port:      fmt.Sprintf(":%d", config.Port),
	}
}

type serverConfig struct {
	Port int `default:"8080"`
}

type Server struct {
	generator *Generator
	port      string
}

func (s Server) Start() {
	http.HandleFunc("/settings", s.handleSettings)

	log.Printf("Starting server on %s\n", s.port)
	log.Fatal(http.ListenAndServe(s.port, nil))
}

type settings struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

func (s Server) handleSettings(w http.ResponseWriter, r *http.Request) {
	var set settings
	if checkError(w, json.NewDecoder(r.Body).Decode(&set)) {
		return
	}

	if checkError(w, s.generator.Set(set.Min, set.Max)) {
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// checkError writes an error response if there was an error and returns true
func checkError(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return true
	}

	return false
}
