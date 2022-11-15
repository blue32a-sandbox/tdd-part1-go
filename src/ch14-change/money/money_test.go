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

func TestSimpleAddition(t *testing.T) {
    five := NewDollar(5)
    sum := five.Plus(NewDollar(5))
    bank := NewBank()
    reduced := bank.Reduce(sum, "USD")
    if !reduced.Equals(NewDollar(10)) {
        t.Fatalf(`sum and ten dollar is not equal. actual: %#v`, sum)
    }
}

func TestPlusReturnsSum(t *testing.T) {
    five := NewDollar(5)
    result := five.Plus(five)
    sum := result.(Sum)
    if !five.Equals(sum.augend) {
        t.Fatalf(`five dollar and result augend is not equal. actual: %#v`, sum.augend)
    }
    if !five.Equals(sum.addend) {
        t.Fatalf(`five dollar and result addend is not equal. actual: %#v`, sum.addend)
    }
}

func TestReduceSum(t *testing.T) {
    sum := Sum{NewDollar(3), NewDollar(4)}
    bank := NewBank()
    result := bank.Reduce(sum, "USD")
    if !result.Equals(NewDollar(7)) {
        t.Fatalf(`result and seven dollar is not equal. actual: %#v`, result)
    }
}

func TestReduceMoney(t *testing.T) {
    bank := NewBank()
    result := bank.Reduce(NewDollar(1), "USD")
    if !result.Equals(NewDollar(1)) {
        t.Fatalf(`result and one dollar is not equal. actual: %#v`, result)
    }
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
    bank := NewBank()
    bank.AddRate("CHF", "USD", 2)
    result := bank.Reduce(NewFranc(2), "USD")
    if !result.Equals(NewDollar(1)) {
        t.Fatalf(`result and one dollar is not equal. actual: %#v`, result)
    }
}

func TestIdentityRate(t *testing.T) {
    bank := NewBank()
    if 1 != bank.Rate("USD", "USD") {
        t.Fatalf("USD to USD rate is not 1")
    }
}
