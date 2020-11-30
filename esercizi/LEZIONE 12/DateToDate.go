package main
import "fmt"
func main(){
  var n,i,sommatoria int
  sommatoria=0
  fmt.Print("Di quale numero vuoi fare la sommatoria?")
  fmt.Scan(&n)
  for i=n;i>0;i--{
    sommatoria = sommatoria + i
  }
  fmt.Print(sommatoria)
}
