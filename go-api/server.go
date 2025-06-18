package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	r := mux.NewRouter()
	r.Use(enableCORS)

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			response := map[string]string{
				"version": "1.0.0",
				"name":    "Stormkit Examples - Go API",
			}

			json.NewEncoder(w).Encode(response)
		}).Methods("GET", "HEAD")

	// This is important, as Stormkit assigns the port dynamically
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
