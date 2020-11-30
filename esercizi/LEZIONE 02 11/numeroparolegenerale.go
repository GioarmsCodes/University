package main
import (
    "fmt"
    "bufio"
    "os"
)
/* Conta il numero delle parole */
func main(){
    var c int
    b:=bufio.NewScanner(os.Stdin)
    for b.Scan(){
        s:= b.Text()
        c=0
        s=" " + s
        for i:=0;i<len(s)-1;i++{
            if s[i] == ' ' && s[i+1] != ' ' {
                c++
            }
        }
        fmt.Println(c)
    }
}
