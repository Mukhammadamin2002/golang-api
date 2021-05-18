package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Title    string `json:"Title"`
	Duration string `json:"duration"`
	Date     string `json:"date"`
}

type Tasks []Task

func allTasks(w http.ResponseWriter, r *http.Request) {
	tasks := Tasks{
		Task{Title: "Housework", Duration: "2 hour", Date: "Monday"},
		Task{Title: "Laisure Activity", Duration: "1 hour", Date: "Tuesday"},
		Task{Title: "Go Suburban Area", Duration: "3:30 hour", Date: "Wednesday"},
		Task{Title: "Holiday", Duration: "whole day", Date: "Friday"},
		Task{Title: "Climbing", Duration: "3 hour", Date: "Thursday"},
		Task{Title: "Helping Other People", Duration: "1:30 hour", Date: "Saturday"},
		Task{Title: "Fishing", Duration: "1:50 hour", Date: "Sunday"},
	}

	fmt.Println("Endpoint Hit: All tasks Endpoint")
	json.NewEncoder(w).Encode(tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tasks", allTasks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
