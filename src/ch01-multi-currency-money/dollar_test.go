package money

import "testing"

func TestMultiplication(t *testing.T) {
    five := Doller{amount: 5}
    five.Times(2)
    if 10 != five.amount {
        t.Fatalf(`amount is not equal. expect: 10, actoual: %d`, five.amount)
    }
}
