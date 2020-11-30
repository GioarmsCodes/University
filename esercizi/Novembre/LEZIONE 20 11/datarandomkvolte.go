package main

import (
    "fmt"
    "math/rand"
    "time"
)
const (
    mesi=12
)
func Isbisestile (a int) (b bool){
    return a%4 == 0 && (a%100 !=0 || a%400 == 0)
}
func GenDate(a int) (g int ,m int){
    mese:=rand.Intn(mesi)

    switch mese {
    case 2:
        if Isbisestile(a){
            return rand.Intn(28)+1,mese+1
        }else{
            return rand.Intn(27)+1,mese+1
        }
    case 1,3, 5,  7, 8, 10, 12:
        return rand.Intn(30)+1,mese+1

    default:
        return rand.Intn(29)+1,mese+1
    }

}
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var anno,volte int
    fmt.Print("Inserisci anno: ")
    fmt.Scan(&anno)
    fmt.Print("Inserisci volte: ")
    fmt.Scan(&volte)
    for i:=0 ; i<volte ; i++{
        x,y:=GenDate(anno)
        fmt.Println(x,y,anno)
    }


}