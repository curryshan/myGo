package simplemath

import "testing"

func TestSqrt1 (t *testing.T) {
    value := Sqrt(4)
    if value != 2 {
        t.Errorf("sqrt(4) is failed. Got %v, expected 2.", value)
    }
}