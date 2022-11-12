package money

import "testing"

func TestEquality(t *testing.T) {
    fiveDollar := NewDollar(5)
    if !fiveDollar.Equals(NewDollar(5)) {
        t.Fatalf("five dollar and five dollar is not equal.")
    }
    if fiveDollar.Equals(NewDollar(6)) {
        t.Fatalf("five dollar and six dollar is equal.")
    }

    fiveFranc := NewFranc(5)
    if fiveFranc.Equals(fiveDollar) {
        t.Fatalf("five franc and five dollar is equal.")
    }
}

func TestCurrency(t *testing.T) {
    dollar := NewDollar(1)
    if dollar.Currency() != "USD" {
        t.Fatalf("dollar currency is not USD")
    }

    franc := NewFranc(1)
    if franc.Currency() != "CHF" {
        t.Fatalf("franc currency is not CHF.")
    }
}

func TestMultiplication(t *testing.T) {
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
