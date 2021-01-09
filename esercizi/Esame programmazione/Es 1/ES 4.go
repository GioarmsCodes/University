package main

import (

    "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
)

type Punto struct {
    etichetta string
    x,y float64

}

type TriangoloRettangolo struct{
    p1,p2,p3 Punto
}

func Distanza(p1 Punto,p2 Punto) float64{
    return math.Sqrt((p1.x - p2.x)*(p1.x - p2.x) + (p1.y - p2.y)*(p1.y - p2.y))
    }

func Stringpunto (p1 Punto) string{
    return fmt.Sprint(p1.etichetta," = (",p1.x,",",p1.y,")")
}

func Max(numero,numero2,numero3 float64) float64{
    var numeri []float64
    var max float64
    numeri = append(numeri,numero)
    numeri = append(numeri,numero2)
    numeri = append(numeri,numero3)
    max = numeri[0]
    for i:=0;i<len(numeri);i++{
        if numeri[i]>max{
            max = numeri[i]
        }
    }
    return max
}
func Min(numero,numero2,numero3 float64) float64{
    var numeri []float64
    var min float64
    numeri = append(numeri,numero)
    numeri = append(numeri,numero2)
    numeri = append(numeri,numero3)
    min = numeri[0]
    for i:=0;i<len(numeri);i++{
        if numeri[i]<min{
            min = numeri[i]
        }
    }
    return min
}
func Istriangolorettangolo(p1 Punto,p2 Punto,p3 Punto) bool{
    var cont,cont2 int
    maxx := Max(p1.x,p2.x,p3.x)
    minx := Min(p1.x,p2.x,p3.x)

    maxy := Max(p1.y,p2.y,p3.y)
    miny := Min(p1.y,p2.y,p3.y)

    for i:=minx;i <= maxx ; i++{
        if i==p1.x {
            cont++
        }
        if i==p2.x {
            cont++
        }
        if i==p3.x {
            cont++
        }
        if cont == 2{
            break
        }
        cont = 0
        }


        for i:=miny;i <= maxy ; i++{
            if i==p1.y {
                cont2++
            }
            if i==p2.y {
                cont2++
            }
            if i==p3.y {
                cont2++
            }
            if cont2 == 2{
                break
            }
            cont2 = 0
            }
    if cont == 2 && cont2 == 2 {
        return true
    }
    return false
}

func AreaMax(t []TriangoloRettangolo) (float64,int){
    var area,areacalc float64
    var index int
    area = 0
    for i:=0 ;i<len(t); i++{
        if Distanza(t[i].p1,t[i].p2) == Max(Distanza(t[i].p1,t[i].p2), Distanza(t[i].p2,t[i].p3) , Distanza(t[i].p1,t[i].p3)){
            areacalc = Distanza(t[i].p2,t[i].p3)*Distanza(t[i].p1,t[i].p3) /2
        }
        if Distanza(t[i].p2,t[i].p3) == Max(Distanza(t[i].p1,t[i].p2), Distanza(t[i].p2,t[i].p3) , Distanza(t[i].p1,t[i].p3)){
            areacalc = Distanza(t[i].p1,t[i].p2)*Distanza(t[i].p1,t[i].p3) /2
        }
        if Distanza(t[i].p1,t[i].p3) == Max(Distanza(t[i].p1,t[i].p2), Distanza(t[i].p2,t[i].p3) , Distanza(t[i].p1,t[i].p3)){
            areacalc = Distanza(t[i].p1,t[i].p2)*Distanza(t[i].p2,t[i].p3) /2
        }
        if areacalc > area {
            area = areacalc
            index = i

        }
    }
    return area,index
}

func main(){
    var flag bool
    var indextemp int
    var x,y float64
    var etichetta string
    var punto Punto
    var punti []Punto
    var triangolo TriangoloRettangolo
    var triangoli []TriangoloRettangolo

    b:= bufio.NewScanner(os.Stdin)

    for b.Scan(){
        s:= b. Text()

        for i,k := range s{
            if k==';' && flag  == false{
                etichetta = s[:i]
                indextemp = i+1
                flag = true
                continue
            }

            if k==';' && flag == true {
                x,_=strconv.ParseFloat(s[indextemp:i],64)
                indextemp = i+1
                flag = false
                continue
            }
            y,_ = strconv.ParseFloat(s[indextemp:],64)
        }
        
        punto.etichetta = etichetta
        punto.x =x
        punto.y =y
        punti = append(punti,punto)
    }

    for i:=0 ;i<len(punti); i++{
        triangolo.p1 = punti[i]
        for j:=i+1;j<len(punti);j++{
            triangolo.p2 = punti[j]
            for m:=j+1 ;m<len(punti);m++{
                triangolo.p3 = punti[m]
                if Istriangolorettangolo(triangolo.p1,triangolo.p2,triangolo.p3){
                    triangoli = append(triangoli,triangolo)
                }
            }
        }
    }
    max,i :=AreaMax(triangoli)

    fmt.Println(Stringpunto(triangoli[i].p1),Stringpunto(triangoli[i].p2),Stringpunto(triangoli[i].p3),max)
}
