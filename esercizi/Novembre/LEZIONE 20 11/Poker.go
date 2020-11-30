package main
//♠	♥ ♦	♣
import (
    "fmt"
    "time"
    "math/rand"
     //"unicode"
     //"strconv"
)
const (
    cart_per_base = 13
    numero_carte = cart_per_base * 4
)
func NtoCarta(n int) (c string) {
    valore:= [13]string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}
    segno:= [4]string{"♠","♥","♦","♣"}
    return valore[n%13] + segno [n/13]
}

func NuovoMazzo (mischia bool) ([numero_carte]string,[numero_carte]int){
    var Mazzo [52]string
    var Mazzonumerico [52]int
    for i:=0 ; i<numero_carte ; i++{
        Mazzonumerico[i]=i
        Mazzo[i]=NtoCarta(i)
    }
    if mischia {
        for i:= numero_carte-1 ; i>0 ; i-- {
            pos := rand.Intn(i)
            Mazzonumerico[pos],Mazzonumerico[i] = Mazzonumerico[i],Mazzonumerico[pos]
            Mazzo[pos],Mazzo[i] = Mazzo[i],Mazzo[pos]
        }
    }
    return Mazzo , Mazzonumerico
}

func coppia(mano []int) (cop bool){
    var contatore int
    slicecopia := mano
    for i:=0 ; i<5 ; i++{
        contatore = 0
        for j:=0 ; j<5 ; j++ {
            if (slicecopia[i]%13) == (mano[j]%13){
                contatore++
                }
            if contatore == 2{
                return true
            }
        }
    }
    return false
}

/* search a on a string */
func main(){
    rand.Seed(time.Now().UTC().UnixNano())
    var carte [52]string
    var cartenumerica [52]int
    var sino int
    var mischia bool
    fmt.Println("P o k e r  ♠ ♥ ♦ ♣ ")
    fmt.Println("Vuoi mischiare le carte? [1 si / 2 no]")
    fmt.Scan(&sino)
    if sino == 1{
        mischia= true
    }else{
        mischia= false
    }
    carte,cartenumerica = NuovoMazzo(mischia)
    fmt.Println(carte)
    fmt.Println(cartenumerica)
    //slice:= cartenumerica[:5]
    slice:=cartenumerica[:5]
    if coppia(slice){
        fmt.Println("Hai una coppia")
    }else{
        fmt.Println("Non hai una coppia")
    }
}
