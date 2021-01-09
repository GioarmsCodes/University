//Programma completo che legge la txt la divina commedia e ti fa fare un po di cose con essa
package main
import (
    "fmt"
    "os"
    "bufio"
    "unicode"
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

func Leggifile2(s string)(lettere map[rune]int,err error){
    lettere = make(map[rune]int)
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
        for _,x:= range l{
            for i:='a';i<='z';i++{
            if i==x{
                lettere[x]++
                break
            }
        }
        }
    }
    return
}

func Leggifile3(s string)(lettere map[string]int,err error){
    var temp string
    lettere = make(map[string]int)
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
        for _,x:= range l{
            if x!= ' '{
                if unicode.IsLetter(x){
                temp +=string(x)
            }
                }else{
                    lettere[temp]++
                    temp = ""
                }
            }
            lettere[temp]++
            temp = ""
        }
        return
    }

func main() {
    var scelta int
    var libro,canto string
    fmt.Printf("Cosa vuoi fare?\n(1)Leggere la divina commedia scegliendo il libro e il canto.\n(2)Contare le lettere lettere dell'alfabeto della divina commedia.\n(3)Contare le parole della divina commedia e ragruppare le parole uguali\nInserisci qui la tua scelta: ")
    fmt.Scan(&scelta)
    switch scelta{
    case 1:
        fmt.Print("Che libro vuoi leggere?[INFERNO/PURGATORIO/PARADISO]: ")
        fmt.Scan(&libro)
        fmt.Print("Che canto vuoi leggere? [METTI IL NUMERO DEL CANTO IN ROMANO]: ")
        fmt.Scan(&canto)
        parole,err := Leggifile("la_divin.txt",libro,canto)
        if err != nil {
            fmt.Println("Errore durante l apertura del file")
        }
        for i:=0 ; i<len(parole) ; i++ {
        fmt.Println(parole[i])
    }
case 2:

    parole,err := Leggifile2("la_divin.txt")
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    for i:='a';i<='z';i++{
        fmt.Println(string(i),": ",parole[i])
    }
case 3:
    var somma int
    parole,err := Leggifile3("la_divin.txt")
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    for k,v := range parole{
        fmt.Println(k,": ",v)
        somma+=v
    }
    fmt.Println("Parole distinte: ",len(parole),"Parole totali: ",somma)
    }

}
