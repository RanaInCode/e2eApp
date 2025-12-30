package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		tasks := []Task{
			{ID: 1, Title: "Learn Docker"},
            {ID: 2, Title: "Master kubernetes"},
			{ID: 3, Title: "Write Go code"},
		}
		json.NewEncoder(w).Encode(tasks)
	})

	http.ListenAndServe(":8080", nil)
}