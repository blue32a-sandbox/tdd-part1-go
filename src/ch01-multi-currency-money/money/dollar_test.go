package money

import "testing"

func TestMultiplication(t *testing.T) {
    five := Dollar{Amount: 5}
    five.Times(2)
    if 10 != five.Amount {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, five.Amount)
    }
}
