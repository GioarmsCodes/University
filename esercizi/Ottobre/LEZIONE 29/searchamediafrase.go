package main
import (
    "fmt"
    "os"
    "bufio"
)
/* search a on a string */
func main(){
    var parole, c float64

    b:=bufio.NewScanner(os.Stdin)
    for b.Scan(){
        s := b.Text()
        parole++
        for i:=0;i<len(s);i++{
            if s[i] == 'a' {
                c++
            }
        }
    }
    fmt.Println("ci sono ",c/parole," a di media")
}
