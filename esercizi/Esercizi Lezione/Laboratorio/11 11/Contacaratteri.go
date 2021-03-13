
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
/* definire un programma che letta una riga da stdin
stampi il numero di occorrenze delle lettere presenti in essa*/
func main() {
	fmt.Println("inserisci una frase")
	b:= bufio.NewScanner(os.Stdin)
	b.Scan()
	s := b.Text()
	s = strings.ToLower(s)
	for c:= 'a';c<='z';c++ {
		var conta int
		for _, r := range s{
			if r==c{
				conta++
			}
		}
		if conta > 0 {
			fmt.Println(string(c),conta)
		}
	}
}
