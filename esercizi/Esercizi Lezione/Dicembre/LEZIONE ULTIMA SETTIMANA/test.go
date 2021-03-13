package main

import "fmt"

func rip(a []int) (b bool) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if j == i {
				continue
			}
			if a[i] == a[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	x := []int{1, 2, 3, 4, 4, 5, 6}
	fmt.Println(rip(x))

}
