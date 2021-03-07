package main

import (

    "fmt"
    "math"
    "unsafe"
)

func memset(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

func main(){
    var A,B,i,j float64
    var zu unsafe.Pointer
    var k int
    var bu unsafe.Pointer
    var z []float64
    var b string
    fmt.Printf("\x1b[2J")
    for ;; {
        memset(bu,32,1760)
        memset(zu,0,7040)

        //b = string(bu)
        for j = 0 ; 6.28>j ; j+=0.07 {
        for i = 0 ; 6.28>i ; i+=0.02 {
            var c,d,e,f,g,h,D,l,m,n,t float64
            c=math.Sin(i)
            d=math.Cos(j)
            e=math.Sin(A)
            f=math.Sin(j)
            g=math.Cos(A)
            h=d+2
            D=1/(c*h*e+f*g+5)
            l=math.Cos(i)
            m=math.Cos(B)
            n=math.Sin(B)
            t=c*h*g-f*e
            var x,y,o,N int
            x=int(40+30*D*(l*h*m-t*n))
            y=int(12+15*D*(l*h*n+t*m))
            o=x+80*y
            N=8*int(((f*e-c*d*g)*m-c*d*e-f*g-l*d*n))
            if  (22>y&&y>0&&x>0&&80>x&&D>z[o]) {
            z[o]=D
            b=".,-~:;=!*#$@"
                fmt.Println(N)
        }
    }
}
    fmt.Printf("\x1b[H")
    for k=0 ; 1761>k ; k++{
    if k%80==0{
        fmt.Print(rune(b[k]))
    }else{
        fmt.Print(rune(10))
    }
    A+=0.04
    B+=0.02

}
    }
    }
