package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct for this handler
type Goodbye struct {
	Log *log.Logger
}

// NewGoodbye function will be exposed and accessed by the servermux
func NewGoodbye(Log *log.Logger) *Goodbye {
	return &Goodbye{Log}
}

func (goodBye *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Goodbye World")
	// fmt.Fprintf(w, "Goodbye World\n")
	w.Write([]byte("Good Bye\n"))
}
