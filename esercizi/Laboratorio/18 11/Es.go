package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Occorrenza(s string) ['z' - 'a' + 1]int {
	var alph ['z' - 'a' + 1]int
	s = strings.ToLower(s)
	for _, x := range s {
		if x >= 'a' && x <= 'z' {
			alph[x-'a']++ //a e' 97 ,la sua posizione in unicode
		}
	}
	return alph
}

func StampaOccorrenza(conta ['z' - 'a' + 1]int) {
	for i := 0; i < len(conta); i++ {
		if conta[i] > 0 {
			fmt.Print(string(i+'a'), ": ", conta[i], "\n")
		}
	}
}
func main() {
	var s string
	fmt.Print("Inserisci valore: ")
	b := bufio.NewScanner(os.Stdin) //Scan della stringa
	b.Scan()
	s = b.Text()

	conta := Occorrenza(s) // Chiamata della funzione per contare le occorrenze
	fmt.Println(conta)
	StampaOccorrenza(conta) // Chiama funzione stampa
}
