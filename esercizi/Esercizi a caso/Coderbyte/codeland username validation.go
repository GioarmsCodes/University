package main
import (
  "fmt"
  "unicode"
)

func CodelandUsernameValidation(str string) string {

  if len(str) >25 || len(str) < 4 {
    return "false"
  }
  if !unicode.IsLetter(rune(str[0])){
    return "false"
  }
   if rune(str[len(str)-1]) == '_'{
    return "false"
  }
  for _,x := range str {
    if unicode.IsLetter(x) || unicode.IsNumber(x) || x=='_'{
      continue
    }
    return "false"
  }
  return "true";

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(CodelandUsernameValidation(readline()))

}
