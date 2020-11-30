package main
import (
	"bufio"
	"fmt"
	"os"
)
/* Stampa la stringa fino a quanto incontra /    */
func main() {
	b := bufio.NewScanner(os.Stdin)
	for b.Scan() {
		s := b.Text()
		for _, x := range s {
			if x == '/' {
				break
			}
			fmt.Print(string(x))
		}
		fmt.Println("")
	}
}
