package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goLambda/internal/handler"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mock Lambda Listener running"))
	})

	mux.HandleFunc("POST /receipts", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
		}
		defer r.Body.Close()

		localHandler := handler.NewHandler()
		localHandler.HandleEndRideLocal(context.TODO(), bodyBytes)
		if err != nil {
			http.Error(w, fmt.Sprintf("Lambda Handler error: %v", err), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "internal server errorr",
				"status":  http.StatusInternalServerError,
			})
		}

		w.WriteHeader(http.StatusOK)
		log.Println("API executed successfully")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "The work has been done successfully",
			"status":  http.StatusOK,
		})
	})

	log.Println("Listening to the port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
