package main
import (
    "fmt"
    "os"
    "bufio"
)


func Leggifile(s string)(parole []string,err error){
    p,e := os.Open(s)
    if e != nil{
        err = e
        return
    }
    defer p.Close()
    b := bufio.NewScanner(p)

    for b.Scan(){
        s := b.Text()
        parole = append(parole,s)
    }
    return
}



func main() {
    var input string
    parole,err := Leggifile(os.Args[1])
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }

    fmt.Println("Inserisci lettere iniziali:")
    fmt.Scan(&input)

    for i:=0 ; i<len(parole) ; i++ {
        if len(parole[i])<len(input){
            continue
        }
    if parole[i][:len(input)] == input {
        fmt.Println(parole[i])
    }
}


}
