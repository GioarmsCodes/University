package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
    "strconv"
)

func moda(serie []float64) float64{
    var mapmoda map[float64]int
    var max int
    var min,dummy float64
    mapmoda = make(map[float64]int)

    for _ ,x := range serie{
        mapmoda[x]++
    }
    for k,v := range mapmoda{
        if v>max{
            max = v
            dummy = k
        }
    }
    min = dummy

    for k,v := range mapmoda{
        if v==max{
            if k<min{
                min = k
            }
        }
    }
    return min
}

func mediana (serie []float64) float64{
    sort.Float64s(serie)
    if len(serie)%2 == 1{
        return serie[(len(serie)-1)/2]
    }
    return (serie[len(serie)/2] + serie[(len(serie)/2) - 1]) / 2

}

func main(){
    var serie []float64
    b := bufio.NewScanner(os.Stdin)

    for b.Scan(){
        num,_ := strconv.ParseFloat(b.Text(),64)
        serie = append(serie,num)

    }

    moda := moda(serie)
    mediana := mediana(serie)
    fmt.Printf("moda: %v\n",moda)
    fmt.Printf("mediana: %v\n",mediana)
}
