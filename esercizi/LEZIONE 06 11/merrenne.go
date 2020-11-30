package main
import "fmt"
func Isprimo(n int) bool {
    var i int
    if n==1{
        return false
        }else{
            for i=2;i < n ; i++ {
                if n%i == 0{
                    break
                }
            }
            if n==i {
                return true
                }else{
                    return false
                }
            }
        }
        func Merrenne(n int) {
            var primo bool
            var c int
            for i:=1 ; c<n ; i*=2{
                primo=Isprimo(i-1)
                if primo==true {
                    fmt.Println(i - 1)
                    c++
                }
            }
        }

        func main(){
            var n int
            fmt.Print("Inserisci numero : ")
            fmt.Scan(&n)
            Merrenne(n)

        }
