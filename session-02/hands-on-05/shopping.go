package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ItemServer struct {
	r *mux.Router
}

// Basically a json server for our purposes
func (s ItemServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(rw, req)
}

var shoppingList []Item

type Item struct {
	Name        string  `json:"name"`
	Supermarket string  `json:"supermarket"`
	Price       float64 `json:"price"`
}

func getItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	for _, item := range shoppingList {
		if item.Name == name {
			enc := json.NewEncoder(w)
			if err := enc.Encode(item); err != nil {
				return
			} else {
				http.Error(w, "Bad item found", http.StatusInternalServerError)
			}
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// encode p to the output.
	enc := json.NewEncoder(w)
	err := enc.Encode(shoppingList)
	if err != nil {
		// if encoding fails, create an error page with code 500.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTotalPrice(w http.ResponseWriter, r *http.Request) {
	total := 0.0
	for _, item := range shoppingList {
		total += item.Price
	}
	enc := json.NewEncoder(w)
	enc.Encode(total)
}

func getItemsForSupermarket(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var items []Item
	for _, item := range shoppingList {
		if item.Supermarket == name {
			items = append(items, item)
		}
	}

	enc := json.NewEncoder(w)
	enc.Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var i Item

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shoppingList = append(shoppingList, i)
}

func removeItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	for i, item := range shoppingList {
		if item.Name == name {
			shoppingList = append(shoppingList[:i], shoppingList[i+1:]...)
			fmt.Fprintf(w, "Item %v removed from shopping list.", name)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func clearShoppingList(w http.ResponseWriter, r *http.Request) {
	shoppingList = nil
	fmt.Fprintf(w, "Shopping list cleared.")
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{name}", getItem).Methods("GET")
	r.HandleFunc("/price", getTotalPrice).Methods("GET")
	r.HandleFunc("/supermarket/{name}", getItemsForSupermarket).Methods("GET")
	r.HandleFunc("/items", addItem).Methods("POST")
	r.HandleFunc("/clear", clearShoppingList).Methods("POST")
	r.HandleFunc("/remove/{name}", removeItem).Methods("POST")

	//wrap gorilla in our ItemServer serving json
	http.Handle("/", &ItemServer{r})
}
