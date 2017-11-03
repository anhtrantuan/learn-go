package math

func Average(xs []float64) float64 {
  length := float64(len(xs))
  if length == 0 { return 0 }

  total := 0.0
  for _, x := range xs {
    total += x
  }
  return total / length
}
