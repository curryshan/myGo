package simplemath

import "testing"

func TestAdd1 (t *testing.T) {
    value := Add(1, 2)
    if value != 3 {
        t.Errorf("add(1, 2) is failed. Got %v, expected 3.", value)
    }

}