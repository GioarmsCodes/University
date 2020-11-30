package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type data struct {
    g, m, a int
}

/* Restituisce la stringa che corrisponde alla data */
func dataStr(d data) string {
  return fmt.Sprintf("%02d/%02d/%04d", d.g, d.m, d.a)
}

/* Converte una stringa (g/m/a) in una data */
func parseData(s string) (d data, ok bool) {
  var err error
  barra1 := strings.IndexRune(s, '/')
  if barra1 < 0 {
    return
  }
  barra2 := strings.IndexRune(s[barra1+1:], '/')
  if barra2 < 0 {
    return
  }
  barra2 += barra1 + 1
  d.g, err = strconv.Atoi(s[:barra1])
  if err != nil {
    return
  }
  d.m, err = strconv.Atoi(s[barra1+1:barra2])
  if err != nil {
    return
  }
  d.a, err = strconv.Atoi(s[barra2+1:])
  if err != nil {
    return
  }
  ok = true
  return
}

/* Determina se l'anno indicato sia bisestile */
func bisestile(a int) bool {
    return a % 4 == 0 && (a % 100 != 0 || a % 400 == 0)
}

/* Restituisce la lunghezza in giorni del mese m nell'anno a */
func lunghezzaMese(m int, a int) int {
  switch m {
  case 11, 4, 6, 9:
    return 30;
  case 2:
    if bisestile(a) {
      return 29
    } else {
      return 28
    }
  default:
    return 31
  }

}

/* Restituisce il numero di giorni compresi fra il 1/1/1970 e la data d.
   Il valore sarà positivo se d è successiva al 1/1/1970, negativo se è precedente */
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

/* Restituisce la distanza in giorni fra le due date */
func distanza(d1, d2 data) int {
  return epoch(d2) - epoch(d1)
}

func main() {
  b := bufio.NewScanner(os.Stdin)
  fmt.Printf("Inserisci la prima data: ")
  b.Scan()
  s := b.Text()
  d1, ok := parseData(s)
  if !ok {
    fmt.Println("Si è verificato un errore")
    return
  }
  fmt.Printf("Inserisci la seconda data: ")
  b.Scan()
  s = b.Text()
  d2, ok := parseData(s)
  if !ok {
    fmt.Println("Si è verificato un errore")
    return
  }
  fmt.Printf("Distanza dall'epoca di %s: %d giorni\n", dataStr(d1), epoch(d1))
  fmt.Printf("Distanza dall'epoca di %s: %d giorni\n", dataStr(d2), epoch(d2))
  fmt.Printf("Distanza %s -> %s: %d giorni\n", dataStr(d1), dataStr(d2), distanza(d1, d2))
}
