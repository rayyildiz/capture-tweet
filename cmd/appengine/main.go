package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("redirecting")
	http.Redirect(w, r, "https://capturetweet.com", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", redirectHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		slog.Info("Defaulting to port", slog.String("port", port))
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
