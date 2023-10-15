package main

import (
	"fmt"
)

func main() {
	var i int = 5
	num := make(chan int)
	num <- i     //sending data to channel num
	num <- i + 5 //sending data to channel num
	fmt.Println(num)

}
