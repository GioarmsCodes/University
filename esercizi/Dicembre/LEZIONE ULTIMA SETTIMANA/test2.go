package main

import "fmt"

func rip(a []int) (b bool) {
	var m map[int]bool
	m = make(map[int]bool)

	for i:=0 ; i<len(a) ;i++{
		_,controllo := m[a[i]]
		if controllo {
			return true
		}
		m[a[i]] = true
	}
	return false
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(rip(x))

}
