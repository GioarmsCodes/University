package main
import "fmt"
/* numero in lettere*/
func main(){
    var numero int
    vocale:= 'a'
    fmt.Scan(&numero)
    unit:= numero%10
    dec:= int(numero/10)
    if numero == 0 {
        fmt.Println("Zero")
        return
    }
    switch dec {
    case 1:
        switch numero {
        case 10:
            fmt.Print("Dieci")
        case 11:
            fmt.Print("undici")
        case 12:
            fmt.Print("dodici")
        case 13:
            fmt.Print("tredici")
        case 14:
            fmt.Print("quattordici")
        case 15:
            fmt.Print("quindici")
        case 16:
            fmt.Print("seidici")
        case 17:
            fmt.Print("diciasette")
        case 18:
            fmt.Print("diciotto")
        case 19:
            fmt.Print("diciannove")
        }
    case 2:
        fmt.Print("Vent")
        vocale= 'i'
    case 3:
        fmt.Print("trent")
    case 4:
        fmt.Print("quarant")
    case 5:
        fmt.Print("cinquant")
    case 6:
        fmt.Print("sesant")
    case 7:
        fmt.Print("settant")
    case 8:
        fmt.Print("ottant")
    case 9:
        fmt.Print("novant")
    }
    if dec == 2 && numero %10 == 0 || numero %10 != 1 || numero %10 != 8 {
        fmt.Print(string(vocale))
    }
    if numero>9 && numero < 20{
        return
    }

    switch unit {
    case 0:
        fmt.Println("")
    case 1:
        fmt.Println("uno")
    case 2:
        fmt.Println("due")
    case 3:
        fmt.Println("tre")
    case 4:
        fmt.Println("quattro")
    case 5:
        fmt.Println("cinque")
    case 6:
        fmt.Println("sei")
    case 7:
        fmt.Println("sette")
    case 8:
        fmt.Println("otto")
    case 9:
        fmt.Println("nove")
    }
}
