package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/endpointB", func(w http.ResponseWriter, r *http.Request) {
		// Take information from service A
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error request A:", err)
			http.Error(w, "Error request A", http.StatusBadRequest)
			return
		}

		// data from serviceA
		fmt.Println("Data from service A:", string(body))

		// response to serviceA
		fmt.Fprintln(w, "Data from service B")
	})

	// Start
	http.ListenAndServe(":8081", nil)
}
