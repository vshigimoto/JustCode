package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/endpointA", func(w http.ResponseWriter, r *http.Request) {
		// data from serviceB
		dataFromB := []byte("Data from service B")
		// data so serviceB
		resp, err := http.Post("http://localhost:8081/endpointB", "application/json", bytes.NewBuffer(dataFromB))
		if err != nil {
			fmt.Println("Error response:", err)
			return
		}
		defer resp.Body.Close()

		// responce to serviceA
		if resp.Status == "200 OK" {
			fmt.Println("Error from service B:", resp.Status)
		} else {
			fmt.Println("Error from service B:", resp.Status)
		}

		fmt.Fprintln(w, "Error from service A")
	})

	// Start
	http.ListenAndServe(":8080", nil)
}
