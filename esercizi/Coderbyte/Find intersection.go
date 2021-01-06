package main
import (
  "fmt"
  "strconv"
  )

func FindIntersection(strArr []string) string {
  var indexspaz int
  var n1,n2 []int
  var s string
  strArr[0] = " "+strArr[0]+","
  strArr[1] = " "+strArr[1]+","

  for i,x := range strArr[0]{
    if x == ' '{
      indexspaz = i
    }
    if x==','{
      temp,_ := strconv.Atoi(strArr[0][indexspaz+1:i])
      n1 = append(n1,temp)
    }
  }
  indexspaz = 0
   for i,x := range strArr[1]{
    if x == ' '{
      indexspaz = i
    }
    if x==','{
      temp,_ := strconv.Atoi(strArr[1][indexspaz+1:i])
      n2 = append(n2,temp)
    }
  }

  for i:=0 ;i<len(n1) ; i++{
    for j:=0 ;j<len(n2) ;j++{
      if n1[i] == n2[j]{
        s +=","+strconv.Itoa(n1[i])
      }
    }
  }
 if len(s) == 0 {
    return "false"
  }
  strArr[0] = s[1:]

  return strArr[0]

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FindIntersection(readline()))

}
