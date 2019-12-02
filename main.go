package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// redirect leads the client through a redirect chain of a specified length.
func redirect(w http.ResponseWriter, req *http.Request) {
	remaining, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/redirect/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		qp := req.URL.Query().Get("n")
		if len(qp) == 0 { // not present
			http.Redirect(w, req, location(remaining-1, remaining), http.StatusTemporaryRedirect)
		} else {
			n, err := strconv.Atoi(qp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else if remaining == 0 {
				w.Write([]byte(endMsg(n)))
			} else {
				http.Redirect(w, req, location(remaining-1, n), http.StatusTemporaryRedirect)
			}
		}
	}
}

func redirectLimit(w http.ResponseWriter, req *http.Request) {
	if strings.Compare(req.URL.Path, "/redirect-limit/") == 0 {
		http.Redirect(w, req, "/redirect-limit/1", http.StatusTemporaryRedirect)
	}
	n, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/redirect-limit/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, req, "/redirect-limit/"+strconv.Itoa(n+1), http.StatusTemporaryRedirect)
	}
}

func location(r int, n int) string {
	return "/redirect/" + strconv.Itoa(r) + "?n=" + strconv.Itoa(n)
}

func endMsg(n int) string {
	if n == 1 {
		return "You followed 1 redirect."
	} else {
		return "You followed " + strconv.Itoa(n) + " redirects."
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: redirect-server <port>")
		os.Exit(1)
	}
	port := os.Args[1]
	http.HandleFunc("/redirect/", redirect)
	http.HandleFunc("/redirect-limit/", redirectLimit)
	http.ListenAndServe(":"+port, nil)
}
