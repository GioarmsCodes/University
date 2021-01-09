//Programma che ti conta le lettere nel programma contando pure i punti le virgole

package main
import (
    "fmt"
    "os"
    "bufio"
    "unicode"
)


func Leggifile(s string,nome string)(string,error){
    var l string
    p,e := os.Open(s)
    if e != nil{
        return "error",e
    }
    defer p.Close()
    b := bufio.NewScanner(p)

    for b.Scan(){
        l = b.Text()
        if len(l) <len(nome){
            continue
        }
        if l[:len(nome)] == nome {
            return l,e
        }
    }
    return "ERROR -- Word not found",e
}



func main() {
    var n string
    b:= bufio.NewScanner(os.Stdin)
    fmt.Println("Dictionary of Oxford , Enter the word you want to search for :  [Exit : Ctrl+C(windows) Ctrl+D(Linux/MacOs)]")
    for b.Scan(){
        n = b.Text()
        n = string(unicode.ToUpper(rune(n[0])))+n[1:]
    parole,err := Leggifile("Oxford English Dictionary.txt",n)
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    fmt.Println(parole)
}
}
