package money

import "testing"

func TestFrancMultiplication(t *testing.T) {
    five := Franc{amount: 5}
    ten := Franc{amount: 10}
    if product := five.Times(2); !ten.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, product.amount)
    }
    fifteen := Franc{amount: 15}
    if product := five.Times(3); !fifteen.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 15, actual: %d`, product.amount)
    }
}
