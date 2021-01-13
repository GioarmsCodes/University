package main

func average(s student) float64 {
  c := 0
  for _, x := range s.Score {
    if x == 32 {
      c += 30
    } else {
      c += x
    }
  }
  return float64(c) / float64(len(s.Score))
}

func average110(s student) float64 {
  m := average(s)
  return 110 * m / 30
}
