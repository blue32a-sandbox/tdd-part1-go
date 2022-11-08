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
