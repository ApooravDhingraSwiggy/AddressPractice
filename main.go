package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Address struct {
	CID   string `json:"cid"`
	ID    string `json:"id"`
	State string `json:"state"`
	Title string `json:"title"`
	Owner *Owner `json:"owner"`
}
type Owner struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var addresses []Address

/*
func getAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}
*/
func getAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range addresses {
		if item.CID == params["cid"] {
			json.NewEncoder(w).Encode(item)
			//return
		}
	}
}

//Delete func
func deleteAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses = append(addresses[:index], addresses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(addresses)
}

func getAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range addresses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //you have to return only if they are equal only one item (no need to send all addresses)
			return
		}
	}
}

func createAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var address Address
	_ = json.NewDecoder(r.Body).Decode(&address)
	address.ID = strconv.Itoa(rand.Intn(100000000))
	addresses = append(addresses, address)
	json.NewEncoder(w).Encode(address)
}

func updateAddress(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//loop over the addresses,range
	//delete the address with the id that you have sent
	//add a new address-the address that we send in the body of postman

	for index, item := range addresses {
		if item.ID == params["id"] {
			addresses = append(addresses[:index], addresses[index+1:]...)
			var address Address
			_ = json.NewDecoder(r.Body).Decode(&address)
			address.ID = params["id"]
			addresses = append(addresses, address)
			json.NewEncoder(w).Encode(address)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	addresses = append(addresses, Address{CID: "C1", ID: "1", State: "Rajasthan", Title: "Address One", Owner: &Owner{Firstname: "John", Lastname: "Doe"}})
	addresses = append(addresses, Address{CID: "C2", ID: "2", State: "Haryana", Title: "Address Two", Owner: &Owner{Firstname: "Steve", Lastname: "Smith"}})
	addresses = append(addresses, Address{CID: "C1", ID: "3", State: "Rajasthan", Title: "Address One.One", Owner: &Owner{Firstname: "John", Lastname: "Doe"}})
	r.HandleFunc("/addresses/{cid}", getAddresses).Methods("GET")
	r.HandleFunc("/addresses/{id}", getAddress).Methods("GET")
	r.HandleFunc("/addresses", createAddress).Methods("POST")
	r.HandleFunc("/addresses/{id}", updateAddress).Methods("PUT")
	r.HandleFunc("/addresses/{id}", deleteAddress).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
