package money

import "testing"

func TestMultiplication(t *testing.T) {
    five := Dollar{amount: 5}
    var product Dollar = five.Times(2)
    if 10 != product.amount {
        t.Fatalf(`amount is not equal. expect: 10, actoual: %d`, product.amount)
    }
    product = five.Times(3)
    if 15 != product.amount {
        t.Fatalf(`amount is not equal. expect: 15, actoual: %d`, product.amount)
    }
}

func TestEquality(t *testing.T) {
    five := Dollar{amount: 5}
    if !five.Equals(Dollar{amount: 5}) {
        t.Fatalf("five dollar and five dollar is not equal.")
    }
    if five.Equals(Dollar{amount: 6}) {
        t.Fatalf("five dollar and six dollar is equal.")
    }
}
