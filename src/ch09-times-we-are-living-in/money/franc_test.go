package money

import "testing"

func TestNewFranc(t *testing.T) {
    five := NewFranc(5)
    var i interface{} = five
    if _, ok := i.(Franc); !ok {
        t.Fatalf("five is not Franc.")
    }
    if five.Amount() != 5 {
        t.Fatalf("the amount is not 5. actual: %d", five.Amount())
    }
}

func TestFrancMultiplication(t *testing.T) {
    five := NewFranc(5)
    ten := NewFranc(10)
    if product := five.Times(2); !ten.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 10, actual: %d`, product.Amount())
    }
    fifteen := NewFranc(15)
    if product := five.Times(3); !fifteen.Equals(product) {
        t.Fatalf(`amount is not equal. expect: 15, actual: %d`, product.Amount())
    }
}

func TestFrancEquality(t *testing.T) {
    five := NewFranc(5)
    if !five.Equals(NewFranc(5)) {
        t.Fatalf("five franc and five franc is not equal.")
    }
    if five.Equals(NewFranc(6)) {
        t.Fatalf("five franc and six franc is equal.")
    }
}
