package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Task struct {
    ID int `json:"id"`
    Title string `json:"title"`
}

var tasks []Task

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "SRE Lab API")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

        http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Connect-Type", "application/json")

		err := json.NewEncoder(w).Encode(tasks)
		if err != nil {
                    http.Error(w, "failed to encode tasks", http.StatusInternalServerError)
		    return
		}
        })

	log.Printf("server started on port %s", port)

	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
