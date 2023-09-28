package main

import (
	"fmt"
)

func main() {
	var balance int = 100

	go func() {
		balance -= 10
	}()

	go func() {
		balance += 10
	}()

	fmt.Println(balance)
}

// package main

// import (
//     "fmt"
// )

// func main() {
//     var counter int = 0

//     go func() {
//         counter++
//     }()

//     go func() {
//         counter++
//     }()

//     fmt.Println(counter)
// }
