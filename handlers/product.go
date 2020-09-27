package handlers

import (
	"log"
	"net/http"

	"github.com/sathishkumar-manogaran/GoLangMicroService/data"
)

type Products struct {
	Log *log.Logger
}

func NewProduct(Log *log.Logger) *Products {
	return &Products{Log}
}

// Decision Making for HTTP method (Get, Post, Put, Delete)
func (product *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// @GetMapping
	if r.Method == http.MethodGet {
		product.getProducts(w, r)
		return
	}

	//catch all other HTTP methods
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// This is the writer; our own json or database value will go to user via ResponseWriter
// @GetMapping
func (product *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	/* products := data.GetProducts()
	jsonProducts, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Error occured while marshalling json", http.StatusInternalServerError)
	}
	w.Write(jsonProducts) */

	// Marshal and Encoder both are same in terms of action
	// But using Encoder we can directly write into io.writer
	// By doing this we can increase performance and also avoid memory allocation to data variable
	// For large size json / data, it is always better to use Encoder
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Error occured while marshalling json", http.StatusInternalServerError)
	}
}
