package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
)

func main(){
    var tempmax,tempmin,voltemax,voltemin,max int
    var temperature map[int]int
    var temp,massimi []int
    temperature = make(map[int]int)
    b:= bufio.NewScanner(os.Stdin)
    b.Split(bufio.ScanWords)
    for b.Scan(){
        temp := b.Text()
        num,_ := strconv.Atoi(temp)
        temperature[num]++
        tempmax = num
        tempmin = num
    }
    for k,v := range temperature{
        if k>=tempmax{
            tempmax = k
            voltemax = v
        }
        if k<=tempmin{
            tempmin = k
            voltemin = v
        }
        temp = append(temp,k)
        if v>max{
            max = v
        }
    }
    fmt.Printf("max : %d, %d volte; min : %d, %d volte\n",tempmax,voltemax,tempmin,voltemin)
    sort.Ints(temp)

    for i:=0 ; i<len(temp) ; i++{
        for k,v := range temperature{
            if temp[i] == k{
                fmt.Printf("%d:%d",k,v)
            }
        }
        if !(i==len(temp)-1){
            fmt.Print(" ")
    }
    fmt.Println("")
    for k,v := range temperature{
        if max == v{
            massimi = append(massimi,k)
        }
    }
    sort.Ints(massimi)
    fmt.Print("temperature ",massimi," con frequenza ",max,", la massima frequenza\n")
}
