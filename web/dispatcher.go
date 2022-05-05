package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	wordApiHostname := os.Getenv("WORD_API_HOSTNAME")
	if wordApiHostname == "" {
		wordApiHostname = "localhost"
	}

	fwd := &forwarder{wordApiHostname, 8080}
	http.Handle("/words/", http.StripPrefix("/words", fwd))
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}

type forwarder struct {
	host string
	port int
}

func (f *forwarder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := net.LookupHost(f.host)
	if err != nil {
		log.Println("Error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	url := fmt.Sprintf("http://%s:%d%s", f.host, f.port, r.URL.Path)
	log.Printf("%s Calling %s", r.URL.Path, url)

	if err = copy(url, w); err != nil {
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
