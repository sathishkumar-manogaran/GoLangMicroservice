package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sathishkumar-manogaran/GoLangMicroService/handlers"
)

func main() {

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

	//env.Parse()

	fmt.Println("Application Started Successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server Up and Running")
		fmt.Fprintf(w, "Server Up and Running\n")
	})

	//1. Start Implementating All the handlers
	log := log.New(os.Stdout, "product-api ", log.LstdFlags)
	helloHandler := handlers.NewHello(log)
	goodByeHandler := handlers.NewGoodbye(log)
	productsHandler := handlers.NewProduct(log)

	//2. Now, register this hello handler in default serve mux
	serverMux := http.NewServeMux()
	serverMux.Handle("/hello", helloHandler)
	serverMux.Handle("/goodbye", goodByeHandler)
	serverMux.Handle("/products", productsHandler)

	// To override default server config
	// This can be used in graceful shutdown too
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serverMux,
		ReadTimeout:  120 * time.Second,
		IdleTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//http.ListenAndServe(":9090", nil)
	// replace the default server mux with above one which is created by us
	// http.ListenAndServe(":9090", serverMux)
	// replace the default server config with custom one
	//server.ListenAndServe()

	// Below all the lines are related to graceful shutdown
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// creating os signal channel to know the exact signal status
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// value of signal channel to log
	signalReceived := <-signalChannel
	log.Println("Received termination request; Singal Received :: ", signalReceived)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)

}
