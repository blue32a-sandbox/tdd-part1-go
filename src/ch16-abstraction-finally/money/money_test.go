package money

import "testing"

func TestEquality(t *testing.T) {
    if NewDollar(5) != NewDollar(5) {
        t.Fatalf("5 dollar and 5 dollar is not equal.")
    }
    if NewDollar(5) == NewDollar(6) {
        t.Fatalf("5 dollar and 6 dollar is equal.")
    }

    if NewFranc(5) == NewDollar(5) {
        t.Fatalf("5 franc and 5 dollar is equal.")
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
    if product := five.Times(2); product != NewDollar(10) {
        t.Fatalf(`product and 10 dollar is not equal. actual: %#v`, product)
    }
    if product := five.Times(3); product != NewDollar(15) {
        t.Fatalf(`product and 15 dollar is not equal. actual: %#v`, product)
    }
}

func TestSimpleAddition(t *testing.T) {
    five := NewDollar(5)
    sum := five.Plus(NewDollar(5))
    bank := NewBank()
    reduced := bank.Reduce(sum, "USD")
    if reduced != NewDollar(10) {
        t.Fatalf(`reduced and 10 dollar is not equal. actual: %#v`, sum)
    }
}

func TestPlusReturnsSum(t *testing.T) {
    five := NewDollar(5)
    result := five.Plus(five)
    sum := result.(Sum)
    if five != sum.augend {
        t.Fatalf(`5 dollar and sum augend is not equal. actual: %#v`, sum.augend)
    }
    if five != sum.addend {
        t.Fatalf(`5 dollar and sum addend is not equal. actual: %#v`, sum.addend)
    }
}

func TestReduceSum(t *testing.T) {
    sum := Sum{NewDollar(3), NewDollar(4)}
    bank := NewBank()
    result := bank.Reduce(sum, "USD")
    if result != NewDollar(7) {
        t.Fatalf(`result and 7 dollar is not equal. actual: %#v`, result)
    }
}

func TestReduceMoney(t *testing.T) {
    bank := NewBank()
    result := bank.Reduce(NewDollar(1), "USD")
    if result != NewDollar(1) {
        t.Fatalf(`result and 1 dollar is not equal. actual: %#v`, result)
    }
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
    bank := NewBank()
    bank.AddRate("CHF", "USD", 2)
    result := bank.Reduce(NewFranc(2), "USD")
    if result != NewDollar(1) {
        t.Fatalf(`result and 1 dollar is not equal. actual: %#v`, result)
    }
}

func TestIdentityRate(t *testing.T) {
    bank := NewBank()
    if 1 != bank.Rate("USD", "USD") {
        t.Fatalf("USD to USD rate is not 1")
    }
}

func TestMixedAddition(t *testing.T) {
    fiveBucks := NewDollar(5)
    tenFrancs := NewFranc(10)
    bank := NewBank()
    bank.AddRate("CHF", "USD", 2)
    result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
    if result != NewDollar(10) {
        t.Fatalf(`result and 10 dollar is not equal. actual: %#v`, result)
    }
}

func TestSumPlusMoney(t *testing.T) {
    fiveBucks := NewDollar(5)
    tenFrancs := NewFranc(10)
    bank := NewBank()
    bank.AddRate("CHF", "USD", 2)
    sum := Sum{fiveBucks, tenFrancs}.Plus(fiveBucks)
    result := bank.Reduce(sum, "USD")
    if NewDollar(15) != result {
        t.Fatalf(`10 dollar and result is not equal. actual: %#v`, result)
    }
}

func TestSumTimes(t *testing.T) {
    fiveBucks := NewDollar(5)
    tenFrancs := NewFranc(10)
    bank := NewBank()
    bank.AddRate("CHF", "USD", 2)
    sum := Sum{fiveBucks, tenFrancs}.Times(2)
    result := bank.Reduce(sum, "USD")
    if NewDollar(20) != result {
        t.Fatalf(`20 dollar and result is not equal. actual: %#v`, result)
    }
}
