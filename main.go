package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sathishkumar-manogaran/GoMicroService/handlers"
)

func main() {
	fmt.Println("Application Started Successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server Up and Running")
		fmt.Fprintf(w, "Server Up and Running\n")
	})

	/* http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Something went wrong"))
			// Alternative
			http.Error(w, "Something went wrong", http.StatusBadRequest)
			return // Need return because http won't return or terminate
		}
		fmt.Fprintf(w, "Request Data :: %s\n", data)
	}) */

	/* http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
		fmt.Fprintf(w, "Goodbye World\n")
	}) */

	//1. Start Implementating All the handlers
	log := log.New(os.Stdout, "product-api ", log.LstdFlags)
	helloHandler := handlers.NewHello(log)
	goodByeHandler := handlers.NewGoodbye(log)

	//2. Now, register this hello handler in default serve mux
	serverMux := http.NewServeMux()
	serverMux.Handle("/hello", helloHandler)
	serverMux.Handle("/goodbye", goodByeHandler)

	//http.ListenAndServe(":9090", nil)
	// replace the default server mux with above one which is created by us
	http.ListenAndServe(":9090", serverMux)
}
