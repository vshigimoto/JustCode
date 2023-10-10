package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/endpointA", func(w http.ResponseWriter, r *http.Request) {
		// data from serviceB
		dataFromB := []byte("Данные от сервиса B")
		// data so serviceB
		resp, err := http.Post("http://localhost:8081/endpointB", "application/json", bytes.NewBuffer(dataFromB))
		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			return
		}
		defer resp.Body.Close()

		// responce to serviceA
		if resp.Status == "200 OK" {
			fmt.Println("Ответ от сервиса B:", resp.Status)
		} else {
			fmt.Println("Ошибка от сервиса B:", resp.Status)
		}

		fmt.Fprintln(w, "Ответ от сервиса A")
	})

	// Start
	http.ListenAndServe(":8080", nil)
}
