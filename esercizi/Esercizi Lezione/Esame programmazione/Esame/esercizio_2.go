package main

import(
  "fmt"
  "os"
)

func isBalanced(sequence string)bool {
  var aperte,chiuse int
  for _, parentesi := range sequence{
    if parentesi == '[' {
      aperte++
    }else{
      chiuse++
    }
    if chiuse > aperte{
        return false
    }
  }
  return chiuse == aperte
}

func main()  {
  var str string
  fmt.Scan(&str)
  for _, char := range str{
    if char != '[' && char != ']'{
      os.Exit(1)
    }
  }

  if isBalanced(str){
      fmt.Println("bilanciata")
    }else{
      fmt.Println("non bilanciata")
    }


    var dictionary map[string]int = make(map[string]int)

for m:=0 ;m< len(str)-1; m++{
    for n:=m+1 ; n<len(str) ; n++{
        newsequence := str[m:n+1]
        if isBalanced(newsequence){
            dictionary[newsequence]++
        }
    }
}


    for i:=len(str);i>0;i--{
        for key,value := range dictionary{
            if len(key) == i{
                fmt.Println(key,value)
            }
        }
    }

}
