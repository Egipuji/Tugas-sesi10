package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Status struct {
	Water       int
	Wind        int
	WaterStatus string
	WindStatus  string
}

func main() {
	http.HandleFunc("/status", handleStatus)

	fmt.Println("server run on 8080")
	http.ListenAndServe(":8080", nil)
}

func getRandomNumb() int {
	return rand.Intn(15)
}

func getStatus(water, wind int) (string, string) {
	waterStatus := "aman"

	if water <= 5 {
		waterStatus = "aman"
	} else if water > 5 && water < 9 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	windStatus := "aman"

	if wind <= 5 {
		windStatus = "aman"
	} else if wind > 5 && water < 9 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	return waterStatus, windStatus
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	wind := getRandomNumb()
	water := getRandomNumb()

	waterStatus, windStatus := getStatus(water, wind)

	status := Status{
		Water:       water,
		Wind:        wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(status)
}
