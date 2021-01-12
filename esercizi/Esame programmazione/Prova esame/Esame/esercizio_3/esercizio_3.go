package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
)
type punto struct{
    etichetta string
    x,y float64
}

func NuovoTragitto() (tragitto []punto){
    var flag bool
    var p punto
    var index int
    b := bufio.NewScanner(os.Stdin)
    for b.Scan(){
        s:= b.Text()

        for i,x := range s{
            if x == ';' && !flag{
                p.etichetta = s[:i]
                flag = true
                index = i+1
                continue
            }
            if x == ';' && flag{
                temp := s[index:i]
                p.x,_= strconv.ParseFloat(temp,64)
                temp = s[i+1:]
                p.y,_= strconv.ParseFloat(temp,64)
                flag = false
                break
            }
        }
        tragitto = append(tragitto , p)
    }
    return
}

func Distanza(a punto,b punto) float64{

    return math.Sqrt((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y))
}

func main(){
    var lunghezza,lunghezza2 float64
    test := NuovoTragitto()
    for i:= 1 ; i<len(test) ; i++{
        lunghezza += Distanza(test[i-1],test[i])
    }
    fmt.Printf("Lunghezza percorso: %.3f\n",lunghezza)

    for i:= 1 ; i<len(test) ; i++{
        lunghezza2 += Distanza(test[i-1],test[i])
        if lunghezza2 > lunghezza/2 {
            fmt.Printf("Punto oltre la metà: %s = (%.1f, %.1f)",test[i].etichetta,test[i].x,test[i].y)
            break
        }
    }

}
