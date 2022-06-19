package handllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func Greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World !")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	Customers := []Customer{
		{Name: "Naveen", City: "Gorakhpur", Zipcode: "273212"},
		{Name: "Arun", City: "Delhi", Zipcode: "265343"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)
	}

}
