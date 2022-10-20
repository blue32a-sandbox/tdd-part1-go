package money

import "testing"

func TestMultiplication(t *testing.T) {
    five := Doller{amount: 5}
    var product Doller = five.Times(2)
    if 10 != product.amount {
        t.Fatalf(`amount is not equal. expect: 10, actoual: %d`, product.amount)
    }
    product = five.Times(3)
    if 15 != product.amount {
        t.Fatalf(`amount is not equal. expect: 15, actoual: %d`, product.amount)
    }
}
