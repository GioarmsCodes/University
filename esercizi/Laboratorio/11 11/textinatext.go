package main
import (
    "fmt"
    "os"
    "bufio"
)
func stringinastring(s string, sottos string) {
    var c int
    c=0
    for i:=0 ; i<len(s) ; i++ {
        if s[i] == sottos[c] {
            c++
            if c==len(sottos)-1 {
                fmt.Print("[",sottos,"] è contenuto nella stringa: [",s,"]")
                fmt.Println()
                return
            }
        }else{
            c=0
        }
    }
    fmt.Print("[",sottos,"] non è contenuto nella stringa: [",s,"]")
    fmt.Println()
}

func main(){
    var s,sottos string
    b:= bufio.NewScanner(os.Stdin)
    fmt.Println("inserisci prima stringa: ")
    b.Scan()
    s=b.Text()
    fmt.Println("inserisci seconda stringa: ")
    b.Scan()
    sottos=b.Text()
    stringinastring(s,sottos)
}
