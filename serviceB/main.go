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
			fmt.Println("Ошибка при чтении данных от сервиса A:", err)
			http.Error(w, "Ошибка при чтении данных от сервиса A", http.StatusBadRequest)
			return
		}

		// data from serviceA
		fmt.Println("Данные от сервиса A:", string(body))

		// response to serviceA
		fmt.Fprintln(w, "Ответ от сервиса B")
	})

	// Start
	http.ListenAndServe(":8081", nil)
}
