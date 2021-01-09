//Programma che prende da riga di comando il nome del libro che vuoi leggere e il numero del canto del vuoi leggere e te lo fa vedere in stdout
package main
import (
    "fmt"
    "os"
    "bufio"
)


func Leggifile(s string,libro string,ncanto string)(parole []string,err error){
    var pass1,pass2,in bool
    p,e := os.Open(s)
    if e != nil{
        err = e
        return
    }
    defer p.Close()
    b := bufio.NewScanner(p)

    for b.Scan(){
        l := b.Text()
        if len(l) == 0{
            continue
        }
        if l == libro {
        parole = append (parole,l)
        pass1 = true
    }
    temp := "CANTO " +ncanto
        if l == temp && pass1{
        pass2 = true
        }

    if pass1 && pass2 && in && l[:5]=="CANTO"{
        return
    }
    if pass1 && pass2 {
        parole = append (parole,l)
        in =true
    }
}
    return
}



func main() {
    parole,err := Leggifile(os.Args[1],os.Args[2],os.Args[3])
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    for i:=0 ; i<len(parole) ; i++ {
    fmt.Println(parole[i])
}

}
