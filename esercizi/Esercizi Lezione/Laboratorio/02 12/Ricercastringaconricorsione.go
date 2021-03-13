package main


import  "fmt"
// ricerca x in a e assume che a sia ordinata

// si considera l elemento mediano di s se x (y)

// se x e' uguale a y ritorna la posizione di y (che e' uguale a lunghezza /2 altrimenti, esempio m=len(a/2))
// se x e' minore di y  effettua la ricerca nella sotto slice a[:m]
// se x e' maggiore di y la ricerca avviene a[m+1:]
// se si arriva a effetuare una ricerca in una slice vuota significa che x non e' presente
func Ricercastringa(a []string , x string) int {

    if len(a)==0{
        return -1
    }

    y:= a[len(a)/2]
    fmt.Println(y)
    if x>y {
        return len(a)/2+Ricercastringa(a[len(a)/2:],x)
    }else if x<y {
        return Ricercastringa(a[:len(a)/2],x)
    }else{
    return len(a)/2
}
}

func main(){
    var x string
    x = "palla"
    a:= []string{"abate","binocolo","cavallo","dinosauro","palla"}
    v:= Ricercastringa(a,x)

    fmt.Println(v)
}
