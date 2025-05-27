package main

import (
	"context"
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
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("lambda executed successfully"))
	})

	log.Println("Listening to the port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
