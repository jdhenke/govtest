package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Starting up.")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		actor := r.URL.Query().Get("actor")
		if actor == "" {
			http.Error(rw, "actor required", http.StatusBadRequest)
			return
		}
		refname := r.URL.Query().Get("refname")
		if refname == "" {
			http.Error(rw, "refname required", http.StatusBadRequest)
			return
		}
		log.Printf("%v initiated check for %v", actor, refname)
		http.Error(rw, http.StatusText(http.StatusOK), http.StatusOK)
	})); err != nil {
		log.Fatal(err)
	}
	log.Println("Shutdown complete.")
}
