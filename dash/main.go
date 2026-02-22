package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// configure the directory name and port
	const resouceDir = "big_buck_bunny"
	port, err := determineListenAddress()
	if err != nil {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(resouceDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", resouceDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
