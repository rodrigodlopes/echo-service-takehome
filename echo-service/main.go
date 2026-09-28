package main

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type Echo struct {
	Headers map[string][]string `json:"Headers"` // Map of request headers.
	Params  map[string][]string `json:"Params"`  // Query parameters.
	Body    string              `json:"Body"`    // The raw request body.
	Path    string              `json:"Path"`    // The request path.
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("failed to read body", "error", err)
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	echo := Echo{
		Headers: r.Header,
		Params:  r.URL.Query(),
		Body:    string(body),
		Path:    r.URL.Path,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(echo); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	slog.Info("listening", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
