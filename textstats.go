package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func parseWords(s string) []string {
	s = strings.ToLower(s)
	return regexp.MustCompile(`\w+`).FindAllString(s, -1)
}

func main() {

	commPortRef := flag.Int("port", 5555, "The port to listen for the data")
	statPortRef := flag.Int("stat", 8080, "The port to dump the statictics")
	flag.Parse()
	commPort := strconv.Itoa(*commPortRef)
	statPort := strconv.Itoa(*statPortRef)

	commMux := http.NewServeMux()
	statMux := http.NewServeMux()

	stat := NewStat()
	queue := make(chan string)

	commMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// "/" matches all URLs so we need to check explicitly
		if r.RequestURI != "/" || r.Method != "POST" {
			http.NotFound(w, r)
		}

		body, _ := ioutil.ReadAll(r.Body)
		for _, word := range parseWords(string(body)) {
			queue <- word
		}
	})

	statMux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		n, err := strconv.Atoi(r.URL.Query().Get("N"))
		if err != nil {
			n = 5
		}

		json, _ := json.Marshal(stat.Dump(n))
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})

	go func() {
		log.Println("Waiting for data at http://localhost:" + commPort)
		log.Fatal(http.ListenAndServe(":"+commPort, commMux))
	}()
	go func() {
		log.Println("Stats at http://localhost:" + statPort)
		log.Fatal(http.ListenAndServe(":"+statPort, statMux))
	}()
	go func() {
		for {
			stat.RecordWord(<-queue)
		}
	}()

	<-make(chan bool) // wait till goroutines do the work
}
