package main
import "fmt"

type data struct{
	gg,mm,aa int
}
// giorno dopo con puntatore
func Giornodopo (g *data) {
	if (*g).gg == 31 && (*g).mm == 12 {
		(*g).gg,(*g).mm=1,1
		(*g).aa++

		}else if (*g).gg < 28 {
			(*g).gg++

			}else{
				switch (*g).mm {
				case 2:
					gbis:= 28
					if Isbisestile((*g).aa) {
						gbis = 29
					}
					if (*g).gg == gbis{
						(*g).gg = 1
						(*g).mm ++
						}else{
							(*g).gg++
						}
					case 1,3,5,7,8,10,12 :
						if (*g).gg == 31 {
							(*g).gg = 1
							(*g).mm ++
							}else{
								(*g).gg++
							}
						default:
							if (*g).gg == 30 {
								(*g).gg = 1
								(*g).mm ++
								}else{
									(*g).gg++
								}
							}
						}
}


func Isbisestile(anno int) bool{
	return anno % 4 == 0 && (anno % 100 != 0 || anno % 400 == 0)
}


func main() {
	var giorno data
	fmt.Print("Inserisci data: esempio [24 11 2001]")
	fmt.Scan(&giorno.gg)
	fmt.Scan(&giorno.mm)
	fmt.Scan(&giorno.aa)
	Giornodopo(&giorno)
	fmt.Println(giorno)
}
