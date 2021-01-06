package main
import "fmt"

func TreeConstructor(strArr []string) string {
  var parents map[string]int
  parents = make(map[string]int)
  var flag bool

   for k:=0 ; k<len(strArr) ; k++{
    for i,x := range strArr[k]{
      if x == ','{
        //child = string(strArr[0][1:i])
        parent := string(strArr[k][i+1:len(strArr[k])-1])
        parents[parent]++
      }
    }
   }

   for _,v := range parents{
     if v>2{
       return "false"
     }
   }

     for _,v := range parents{
     if v == 2{
       flag = true
     }
   }
   if !flag {
     return "false"
   }

  return "true"

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(TreeConstructor(readline()))

}
