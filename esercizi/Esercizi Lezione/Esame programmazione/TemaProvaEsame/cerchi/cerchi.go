package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
)

type Cerchio struct{
    nome string
    x,y,raggio float64
}

func newCircle(descr string) Cerchio{
    var cerc Cerchio
    dol := strings.Split(descr," ")
     cerc.nome = dol[0]
     cerc.x,_ = strconv.ParseFloat(dol[1],64)
     cerc.y,_ = strconv.ParseFloat(dol[2],64)
     cerc.raggio,_ = strconv.ParseFloat(dol[3],64)
     return cerc
}

func distanzaPunti(x1,y1,x2,y2 float64) float64{
    return math.Sqrt((x1 - x2)*(x1 - x2) + (y1 - y2)*(y1 - y2))
}

func tocca(cerc1, cerc2 Cerchio) bool{
    if math.Abs(cerc1.raggio + cerc2.raggio - distanzaPunti(cerc1.x,cerc1.y,cerc2.x,cerc2.y)) < 1e-6  || math.Abs(cerc1.raggio - cerc2.raggio - distanzaPunti(cerc1.x,cerc1.y,cerc2.x,cerc2.y)) < 1e-6 {
        return true
    }
    return false
}

func maggiore(cerc1, cerc2 Cerchio) bool{
    if cerc1.raggio > cerc2.raggio{
        return true
    }
    return false
}

func main(){
    var prevcerc Cerchio
    prevcerc.nome = ""
    prevcerc.x = 0
    prevcerc.y = 0
    prevcerc.raggio = 0
    b := bufio.NewScanner(os.Stdin)
    for b.Scan(){
        s := b.Text()
        c := newCircle(s)
        fmt.Printf("{%s %v %v %v} ",c.nome,c.x,c.y,c.raggio)
        if tocca(c,prevcerc){
            fmt.Print("tangente, ")
        }else{
            fmt.Print("non tangente, ")
        }
        if maggiore(c,prevcerc){
            fmt.Print("maggiore")
        }else{
            fmt.Print("minore o uguale")
        }
        fmt.Println("")
        prevcerc = c
    }


}
