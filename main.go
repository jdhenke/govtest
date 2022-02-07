package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Starting up.")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := approve(r.URL.Query()); err != nil {
			log.Println(err)
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
		http.Error(rw, http.StatusText(http.StatusOK), http.StatusOK)
	})); err != nil {
		log.Fatal(err)
	}
	log.Println("Shutdown complete.")
}

func approve(q url.Values) error {
	actor := q.Get("actor")
	if actor == "" {
		return fmt.Errorf("no actor provided")
	}
	refname := q.Get("refname")
	if refname == "" {
		return fmt.Errorf("no refname provided")
	}
	sha := q.Get("sha")
	if sha == "" {
		return fmt.Errorf("no sha provided")
	}
	resp, err := http.Get(fmt.Sprintf("https://github.com/jdhenke/govtest/commit/%s.diff", sha))
	if err != nil {
		return fmt.Errorf("getting sha diff %s: %v", sha, err)
	}
	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return fmt.Errorf("reading sha diff %s: %v", sha, err)
	}
	const max = 1024
	if n > max {
		return fmt.Errorf("diff larger than max: %v > %v", n, max)
	}
	log.Printf("%s %s by %s approved.", sha, refname, actor)
	return nil
}
