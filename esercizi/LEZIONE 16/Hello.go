package main
import "fmt"
func main(){
  var n,i int
  fmt.Print("quante volte vuoi essere salutato ,padrone?")
  fmt.Scan(&n)
  for i=0;i<n;i++{
    fmt.Println("ciao")
  }
}
