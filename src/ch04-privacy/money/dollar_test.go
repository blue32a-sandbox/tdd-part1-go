package money

import "testing"

func TestNewDollar(t *testing.T) {
    five := NewDollar(5)
    var i interface{} = five
    if _, ok := i.(Dollar); !ok {
        t.Fatalf("five is not Dollar.")
    }
    if five.amount != 5 {
        t.Fatalf("the amount is not 5. actual: %d", five.amount)
    }
}

func TestMultiplication(t *testing.T) {
    five := Dollar{amount: 5}
    ten := Dollar{amount: 10}
    if product := five.Times(2); !ten.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, product.amount)
    }
    fifteen := Dollar{amount: 15}
    if product := five.Times(3); !fifteen.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 15, actual: %d`, product.amount)
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
