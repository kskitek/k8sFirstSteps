package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"k8sFirstSteps"
)

const defaultPort = "8080"

var decoder = k8sFirstSteps.NewDecoder()

func main() {
	go startPinging()

	http.HandleFunc("/data", handler)
	port := getPort()
	log.Println("Listening on port: " + port)
	panic(http.ListenAndServe(":"+port, nil))
}

func startPinging() {
	url := os.Getenv("BLUE_URL")
	client := http.Client{
		Timeout: time.Second * 5,
	}
	request, err := http.NewRequest(http.MethodPost, url+"/generate", nil)
	if err != nil {
		panic(err)
	}

	sleep, err := time.ParseDuration(os.Getenv("SLEEP_DURATION"))
	if err != nil {
		panic(err)
	}
	for {
		<-time.After(sleep)
		resp, err := client.Do(request)
		if err != nil {
			log.Println("Failed request to BLUE: " + err.Error())
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response: {code: %d, body: %s}\n", resp.StatusCode, body)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data k8sFirstSteps.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	decodedValue, err := decoder.Decode(data.Payload)
	if err != nil {
		log.Println(err, decodedValue)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("GOT: [%s] at %s", decodedValue, data.GeneratedAt.Format(time.RFC3339))
	// TODO save it to file
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}
