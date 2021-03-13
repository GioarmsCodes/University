// qual'è l'output del programma?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Println("digita una riga")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(len(s)) //  lunghezza stringa
	fmt.Println(xxx(s)) // output del stringa togliendo le ripetizioni
}

func xxx(s string) string {
    var out string
	for i := 0; i < len(s); i++ {
		if strings.Index(out, string(s[i])) == -1 {
			out += string(s[i])
		}
	}
	return out
}
