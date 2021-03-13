package main
import (
    "fmt"
    "os"
    "strconv"
)
type data struct {
    g,m,a int
}

func lunghezzaMese(m int,a int) (x int){
        switch m {
        case 2:
            if bisestile(a){
                return 29
            }else{
                return 28
            }
        case 1,3, 5,  7, 8, 10, 12:
            return 31
        default:
            return 30
        }
}
func bisestile(a int) (b bool){
    return a % 4 == 0 && (a % 100 != 0 || a % 400 == 0)
}

func epoch(d data) (g int) {
  if d.a >= 1970 {
    for a := 1970; a < d.a; a++ {
      g += 365
      if bisestile(a) {
        g++
      }
    }
    for m := 1; m < d.m; m++ {
      g += lunghezzaMese(m, d.a)
    }
    g += d.g - 1
    return
  } else {
    for a := d.a + 1; a < 1970; a++ {
      g -= 365
      if bisestile(a) {
        g--
      }
    }
    for m := d.m; m < 13; m++ {
      g -= lunghezzaMese(m, d.a)
    }
    g += d.g
    return
  }
}

func main() {
    var data1,data2 data
    var distanza,t,c,b int
    c=0
    for i,x:= range os.Args[1]{
        //giorno
        if x=='/' && c==0 {
            t,_=strconv.Atoi(os.Args[1][:i])
            data1.g = t
            c=1
            b=i+1
            continue
        }
        //mese
        if x=='/' && c==1{
            t,_=strconv.Atoi(os.Args[1][b:i])
            data1.m = t
            b=i+1
            //anno
            t,_=strconv.Atoi(os.Args[1][b:])
            data1.a = t

        }
    }

        b=0
        c=0
    for i,x:= range os.Args[2]{
        //giorno
        if x=='/' && c==0 {
            t,_=strconv.Atoi(os.Args[2][:i])
            data2.g = t
            c=1
            b=i+1
            continue
        }
        //mese
        if x=='/' && c==1{
            t,_=strconv.Atoi(os.Args[2][b:i])
            data2.m = t
            b=i+1
            //anno
            t,_=strconv.Atoi(os.Args[2][b:])
            data2.a = t

        }
    }


    distanza = epoch(data1) - epoch (data2)

    fmt.Println(data1)
    fmt.Println(data2)
    if distanza > 0 {
        fmt.Println(distanza)
    }else if distanza <0{
        fmt.Println(-distanza)
    }else{
        fmt.Println(distanza,"sono lo stesso giorno")
    }


    }
