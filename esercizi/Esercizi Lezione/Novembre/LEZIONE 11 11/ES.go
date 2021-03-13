package main
import "fmt"
type data struct{
    g,m,a int
}
func datetodate(anno int) (int,int,int){
    var giorno,mese int
    if anno%400==0{
        giorno=29
        mese=2
        return giorno,mese,anno
    }else if anno%4==0 && anno%100!=0{
        giorno=29
        mese=2
        return giorno,mese,anno
    }else{
        giorno=28
        mese=2
        return giorno,mese,anno
    }
}
func main(){
    var d data
    var n int
    fmt.Print("Inserisci anno: ")
    fmt.Scan(&n)
    d.g , d.m , d.a = datetodate(n)
    fmt.Println(d)

}
