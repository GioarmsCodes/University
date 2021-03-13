package main

import "fmt"
  // Programma che converte numeri interi positivi da base 10 a base [2-16]
func Basechange(numero int, base int) (string, bool) {

	const cifre = "0123456789ABCDEF"
	var tradotto string
	var error bool
	error = false
	if numero < 0 || base > 16 {
		error = true
		return tradotto, error
	}
	for {
		tradotto = string(cifre[numero%base]) + tradotto
		numero /= base
		if numero == 0 {
			break
		}
	}
	return tradotto, error
}
func main() {
	var n, b int
	var convertito string
	var e bool
	fmt.Println("// Programma che converte numeri interi positivi da base 10 a base [2-16] // ")
	fmt.Print("Inserisci numero: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci base: ")
	fmt.Scan(&b)
	convertito, e = Basechange(n, b)
	if !e {
		fmt.Println(convertito)
	} else {
		fmt.Println("Hai inserito dei numeri non idonei")
	}

}
