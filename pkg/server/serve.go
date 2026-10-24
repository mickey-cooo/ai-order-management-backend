package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Serve runs a bare HTTP server with /healthz on $PORT, falling back to defaultPort.
func Serve(name, defaultPort string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	log.Printf("%s service listening on :%s", name, port)
	srv := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Fatal(srv.ListenAndServe())
}
