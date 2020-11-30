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
	fmt.Println(xxx(s)) // output del stringa togliendo le ripetizioni
}

func xxx(s string) string {
    var out string
	for _,x:=range s {
		if strings.Index(out, string(x)) == -1 {
			out += string(x)
		}
	}
	return out
}
