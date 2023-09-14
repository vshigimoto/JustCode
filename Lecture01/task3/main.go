package main

import "fmt"

func compareTwoSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		found := false
		for _, n := range b {
			if v == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3}
	b := []int{3, 2, 3}
	fmt.Println(compareTwoSlices(a, b))
}
