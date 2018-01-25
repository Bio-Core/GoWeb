package main

import (
	"log"
	"net/http"
	"time"
)

// Log files should typically be located in a rotating log file on the filesystem.
// Investigate the possibility of having this in a /var/log (on Linux) directory
// that rotates daily for a week.

//Logger is a log that prints to the console when web pages are hit
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		if goTest == false {
			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		}
	})
}
