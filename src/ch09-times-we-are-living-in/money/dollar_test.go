package money

import "testing"

func TestNewDollar(t *testing.T) {
    five := NewDollar(5)
    var i interface{} = five
    if _, ok := i.(Dollar); !ok {
        t.Fatalf("five is not Dollar.")
    }
    if five.Amount() != 5 {
        t.Fatalf("the amount is not 5. actual: %d", five.Amount())
    }
}

func TestDollarMultiplication(t *testing.T) {
    five := NewDollar(5)
    ten := NewDollar(10)
    if product := five.Times(2); !ten.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, product.Amount())
    }
    fifteen := NewDollar(15)
    if product := five.Times(3); !fifteen.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 15, actual: %d`, product.Amount())
    }
}

func TestDollarEquality(t *testing.T) {
    five := NewDollar(5)
    if !five.Equals(NewDollar(5)) {
        t.Fatalf("five dollar and five dollar is not equal.")
    }
    if five.Equals(NewDollar(6)) {
        t.Fatalf("five dollar and six dollar is equal.")
    }
}
