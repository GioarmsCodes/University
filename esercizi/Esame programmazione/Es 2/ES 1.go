package main

import (
    "fmt"
    "os"
    "unicode"
    "bufio"


)

func main(){
    var numerocar, flag bool

    var contmai , contmin ,contndec , notcl int
    b := bufio.NewScanner(os.Stdin)

    for b.Scan(){
        s := b.Text()
        if len (s)>= 12 {
            numerocar = true
        }

        for _,x := range s {
            if unicode.IsDigit(x) {
                contndec++
                flag = true
            }
            if unicode.IsLetter(x){
                for i:= 'a' ; i<'z' ; i++{
                    if i == x {
                        contmin++
                        flag = true
                    }
                }
                for i:='A' ; i<='Z' ; i++{
                    if i == x{
                        contmai++
                        flag = true
                    }
                }
}
            if !flag {
                notcl++
            }
            flag = false
        }

        if numerocar && contndec >= 3 && contmin >=2 && contmai >=2 && notcl >=4 {
            fmt.Println("La pw è ben definita!")
        }else{
            fmt.Println("La pw non e' definita correttamente")
            if !numerocar {
                fmt.Println("La pw deve avere una lunghezza minima di 12 caratteri")
            }
            if contndec < 3 {
                fmt.Println("Almeno 3 caratteri della pw devono rappresentare delle cifre decimali")
            }
            if contmin < 2 {
                fmt.Println("Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
            }
            if contmai < 2 {
                fmt.Println("Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole")
            }
            if notcl < 4 {
                fmt.Println(" Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali")
            }
        }
        numerocar = false
        contndec = 0
        contmin =0
        contmai =0
        notcl =0
        }

        }
