package main
import "fmt"

func BracketMatcher(str string) string {

  var b1 ,b2 int

  for _,x := range str {
    if x == '('{
      b1++
    }

    if x == ')'{
      b2++
    }

    if b2>b1{
      return "0"
    }
  }
  if b1 != b2 {
    str = "0"
  }else{
    str = "1"
  }

  return str;

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(BracketMatcher(readline()))

}
