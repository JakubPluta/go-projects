package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Departure struct {
	Departing    string `json:"departing"`
	Destination  string `json:"destination"`
	FlightNumber string `json:"flightNumber"`
	Gate         string `json:"gate"`
	Status       string `json:"status"`
}

type Response map[string]Departure

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getFlights(w http.ResponseWriter, r *http.Request) {
	cassandraToken := goDotEnvVariable("DATASTAX_CASSANDRA_TOKEN")
	baseURL := goDotEnvVariable("DATASTAX_ENDPOINT_URI")
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Cassandra-Token", cassandraToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	json.NewEncoder(w).Encode(responseObject)

}

func main() {

	http.HandleFunc("/flights", getFlights)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
