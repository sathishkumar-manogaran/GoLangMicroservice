package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct
type Hello struct {
	Log *log.Logger
}

// NewHello function will be exposed and accessed by the servermux
func NewHello(Log *log.Logger) *Hello {
	return &Hello{Log}
}

// ServeHTTP is the default interface in GO, we are implementing that interface like below
func (hello *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello.Log.Println("Hello World")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Something went wrong"))
		// Alternative
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return // Need return because http won't return or terminate
	}
	fmt.Fprintf(w, "Request Data :: %s\n", data)
}
