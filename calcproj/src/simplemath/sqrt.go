package simplemath

import "math"

func Sqrt (v int) int {
    value := math.Sqrt(float64(v))
    return int(value)
}