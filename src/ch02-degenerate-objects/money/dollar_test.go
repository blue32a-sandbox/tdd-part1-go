package money

import "testing"

func TestMultiplication(t *testing.T) {
    five := Dollar{Amount: 5}
    var product Dollar = five.Times(2)
    if 10 != product.Amount {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, product.Amount)
    }
    product = five.Times(3)
    if 15 != product.Amount {
        t.Fatalf(`amount is not equal. expect: 15, actual: %d`, product.Amount)
    }
}
