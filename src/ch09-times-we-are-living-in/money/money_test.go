package money

import "testing"

func TestMoneyEquality(t *testing.T) {
    fiveFranc := NewFranc(5)
    fiveDollar := NewDollar(5)
    if fiveFranc.Equals(fiveDollar) {
        t.Fatalf("five franc and five dollar is equal.")
    }
    if fiveDollar.Equals(fiveFranc) {
        t.Fatalf("five dollar and five franc is equal.")
    }
}

func TestCurrency(t *testing.T) {
    dollar := NewDollar(1)
    if dollar.Currency() != "USD" {
        t.Fatalf("dollar currency is not %s.", "USD")
    }

    franc := NewFranc(1)
    if franc.Currency() != "CHF" {
        t.Fatalf("franc currency is not %s.", "CHF")
    }
}
