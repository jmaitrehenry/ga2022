package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	wordApiPortURL := os.Getenv("WORD_API_URL")
	if wordApiPortURL == "" {
		wordApiPortURL = "http://words"
	}

	fwd := &forwarder{wordApiPortURL}
	http.Handle("/words/", http.StripPrefix("/words", fwd))
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}

type forwarder struct {
	url string
}

func (f *forwarder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s%s", f.url, r.URL.Path)
	log.Printf("%s Calling %s", r.URL.Path, url)

	if err := copy(url, w); err != nil {
		log.Println("Error", err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func copy(url string, w http.ResponseWriter) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	for header, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	return err
}
