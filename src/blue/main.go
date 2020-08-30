package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"k8sFirstSteps"
)

const defaultPort = "8081"

var (
	encoder = k8sFirstSteps.NewEncoder()

	counter int64

	url    = os.Getenv("RED_URL")
	client = http.Client{Timeout: time.Second * 5}
)

func main() {
	http.HandleFunc("/generate", handler)
	port := getPort()
	log.Println("Listening on port: " + port)
	panic(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Handling request")

	payload := fmt.Sprintf("request number: %d", counter)
	atomic.AddInt64(&counter, 1)

	var err error
	payload, err = encoder.Encode(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	data := k8sFirstSteps.Data{
		Payload:     payload,
		GeneratedAt: time.Now(),
	}
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	request, err := http.NewRequest(http.MethodPost, url+"/data", body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	log.Println("Sent data to RED")
	w.WriteHeader(http.StatusAccepted)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}
