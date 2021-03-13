package main
 // kevin passa continuo al parziale 
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type frequenze ['z' - 'a' + 1]int // kevin fai schifo

func Occorrenza(s string,alph *frequenze) {
	s = strings.ToLower(s)
	for _, x := range s {
		if x >= 'a' && x <= 'z' {
			(*alph)[x-'a']++ //a e' 97 ,la sua posizione in unicode
		}
	}
}

func StampaOccorrenza(conta ['z' - 'a' + 1]int) {
	for i := 0; i < len(conta); i++ {
		if conta[i] > 0 {
			fmt.Print(string(i+'a'), ": ", conta[i], "\n")
		}
	}
}
// test armani
func Riempialfabeto() (list [26]string){
	for i:= 'a' ; i<='z' ; i++ {
		list[i-'a']=string(i)
	}
	return
}
// test armani
func main() {
	var alfabeto [26]string
	var alph frequenze
	var s string
	fmt.Print("Inserisci valore: [esci per uscire]")
	b := bufio.NewScanner(os.Stdin) //Scan della stringa
	for b.Scan(){
		s = b.Text()
		if s == "-1" {
			break
		}
		//TEST armani
		alfabeto = Riempialfabeto()
		fmt.Println(alfabeto)
		//TEST armani

		Occorrenza(s,&alph) // Chiamata della funzione per contare le occorrenze
		fmt.Println(alph)
		StampaOccorrenza(alph) // Chiama funzione stampa

	}
}
