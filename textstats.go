package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func parseWords(s string) []string {
	s = strings.ToLower(s)
	// `\pL` matches any Unicode character (https://github.com/google/re2/wiki/Syntax)
	return regexp.MustCompile(`\pL+`).FindAllString(s, -1)
}

func main() {

	var commPortFlag, statPortFlag int
	var isDebug bool
	flag.IntVar(&commPortFlag, "port", 5555, "The port to listen for the data")
	flag.IntVar(&statPortFlag, "stat", 8080, "The port to dump the statictics")
	flag.BoolVar(&isDebug, "debug", false, "Turn on additional logging")
	flag.Parse()

	commPort := strconv.Itoa(commPortFlag)
	statPort := strconv.Itoa(statPortFlag)

	stat := NewStat()
	queue := make(chan string)

	// Serving stats
	go func() {
		http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			n, err := strconv.Atoi(r.URL.Query().Get("N"))
			if err != nil {
				n = 5
			}
			json, _ := json.Marshal(stat.Dump(n))
			w.Header().Add("Content-Type", "application/json")
			w.Write(json)
		})
		log.Println("Stats at http://localhost:" + statPort + "/stats")
		log.Fatal(http.ListenAndServe(":"+statPort, nil))
	}()

	// Receiveing the data
	go func() {
		log.Println("Listening TCP localhost:" + commPort)
		ln, err := net.Listen("tcp", ":"+commPort)
		if err != nil {
			panic(err)
		}
		for {
			conn, _ := ln.Accept()
			go func(conn net.Conn) {
				connbuf := bufio.NewReader(conn)
				defer conn.Close()
				for {
					str, err := connbuf.ReadString('\n')
					if err != nil {
						break
					}
					words := parseWords(str)
					for _, word := range words {
						queue <- word
					}
					if isDebug {
						log.Println("Received " + strconv.Itoa(len(words)) + " words from " +
							conn.RemoteAddr().String())
					}
				}
			}(conn)
		}
	}()

	// Synchronization
	for {
		stat.RecordWord(<-queue)
	}
}
