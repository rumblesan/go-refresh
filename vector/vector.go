package vector

import (
	"math"
	"math/rand"
)

type Vector struct {
  X float64 `json:"x"`
  Y float64 `json:"y"`
}

func RandomVector() *Vector {
  x := rand.Float64()
  y := rand.Float64()
  return &Vector{X: x, Y: y}
}

func (v Vector) Magnitude() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
