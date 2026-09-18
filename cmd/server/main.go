// Command server runs the githubbox HTTP server. Vercel's Go framework preset
// builds this entrypoint and starts it with PORT set.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dferber90/githubbox/internal/githubbox"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", githubbox.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
