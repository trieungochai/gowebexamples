// This example will show how to encode and decode JSON data using the encoding/json package
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		kai := User{
			FirstName: "Kai",
			LastName:  "T10H",
			Age:       69,
		}

		json.NewEncoder(w).Encode(kai)
	})

	http.ListenAndServe(":8080", nil)
}

// $ go run json.go

// $ curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode
// Elon Musk is 48 years old!

// $ curl -s http://localhost:8080/encode
// {"firstname":"John","lastname":"Doe","age":25}
