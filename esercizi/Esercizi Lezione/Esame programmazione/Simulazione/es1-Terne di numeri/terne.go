package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "bufio"
)

func StampTerne(n []int){
  //fmt.Println(len(n))
    for i:=0 ;i<len(n)-2 ;i++{
      //fmt.Println("i: ",i)
        for j:=i+1 ;j<len(n)-1 ; j++{
          //fmt.Println("j: ",j)
            for m:=j+1 ;m<len(n) ; m++{
              //fmt.Println("m: ",m)

                if n[i]<n[j] && n[j]<n[m]{
                    fmt.Println(n[i],"<",n[j],"<",n[m])
                }
            }
        }
    }
}

func main(){

    var myArr []int
rr := bufio.NewReader(os.Stdin)
linea, _ := rr.ReadString('\n')
nums := strings.Split(linea[:len(linea)-2], " ")
for i:=0; i<len(nums); i++ {
x, _ := strconv.Atoi(nums[i])
if x > 0 {
myArr = append(myArr, x)
}
}
fmt.Print("Valori inseriti:\n", myArr,"\n")

StampTerne(myArr)

}
