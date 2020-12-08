package main

import (

    "fmt"
    "math/rand"
    "time"
)

type aula [][]int

func CreaAula(riga int,colonna int) (aula,bool){
    var error bool
    error = true
    if riga <1 || colonna <1 {
        error = false
    }
    a := make(aula, riga)

    for i := range a {
        a[i] = make([]int, colonna)

    }
    if !error {
        a = RandAula()
    }
    return a,error
}

func RandAula() aula{

    var riga,colonna int
    riga = rand.Intn(10)+1
    colonna = rand.Intn(10)+1
    a,_ := CreaAula(riga,colonna)
    return a
}

func StampaAula(a aula){
    for i:=0 ;i<len(a);i++{
        for j:=0 ; j<len(a[i]);j++{
            if a[i][j] == 0 {
                fmt.Print("_ ")
            }else{
                fmt.Print("X ")
            }
        }
        fmt.Println("")
    }
}

func Occupa(a aula, fila int, col int) bool {
    var error bool
    error = true
    for i:=0 ; i<len(a) ; i++ {
        for j:=0 ; j<len(a[i]) ; j++ {
            if i == fila && j == col {
                if a[i][j] != 0 {
                    error = false
                }else{
                    a[i][j] = 1
                }
            }
        }
    }
    return error
}

func Libera(a aula, fila int, col int) bool {
    var error bool
    error = true
    for i:=0 ; i<len(a) ; i++ {
        for j:=0 ; j<len(a[i]) ; j++ {
            if i == fila && j == col {
                if a[i][j] != 1 {
                    error = false
                }else{
                    a[i][j] = 0
                }
            }
        }
    }
    return error
}

func VerificaDistanziamento(a aula, fila int, col int) (error bool) {
    error = true
    if fila > 0 && fila < len(a) && col > 0 && col< len(a[0]){
        if a[fila][col+1] == 1 || a[fila][col-1] == 1 || a[fila+1][col] == 1 || a[fila+1][col] == 1 {
            error = false
            return
        }else{
            error = true
            return
        }
    }
    if fila == 0 {
        if a[fila+1][col] == 1 || a[fila][col+1] == 1 || a[fila][col-1] == 1{
            error = false
            return
        }else{
            error = true
            return
    }
}
    if col == 0{
        if a[fila][col+1] == 1 || a[fila+1][col] == 1 || a[fila-1][col] == 1{
            error = false
            return
        }else{
            error = true
            return
        }
    }
    if fila == len(a)-1 {
        if a[fila-1][col] == 1 || a[fila][col+1] == 1 || a[fila][col-1] == 1{
            error = false
            return
        }else{
            error = true
            return
    }
}
    if col == len(a[0])-1{
        if a[fila][col-1] == 1 || a[fila+1][col] == 1 || a[fila+1][col] == 1{
            error = false
            return
        }else{
            error = true
            return
    }
}
 return
}
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var riga,colonna int
    var test aula
    var prova,prova2,prova3,prova4 bool
    fmt.Print("Inserisci Righe: ")
    fmt.Scan(&riga)
    fmt.Print("Inserisci Banchi: ")
    fmt.Scan(&colonna)

    //Creo una aula con le grandezze che voglio io
    test,prova = CreaAula(riga,colonna)
    StampaAula(test,prova) //Se prova = false ho generato una aula random perche' i valori inseriti erano errati

    fmt.Println()
    fmt.Println()

    //Provo il funzionamento della funzione Ocuppa che dato una aula e una posizione occupo tale posizione
    prova2 = Occupa(test ,3,2)
    prova2 = Occupa(test ,0,0)
    StampaAula(test)

    //Provo il funzionamento della funzione VerificaDistanziamento Che verifica il distanziamento sociale
    prova4 = VerificaDistanziamento(test,1,0)
    fmt.Println(prova4)


    fmt.Println()
    fmt.Println()
    //Provo il funzionamento della funzione Libera che e' il contrario di occupa 
    prova3 = Libera(test ,3,2)
    StampaAula(test)
    fmt.Println(prova,prova2,prova3)


}
