package money

import "reflect"

type Franc struct {
    Money
}

func NewFranc(amount int64) Franc {
    return Franc{Money{amount: amount}}
}

func (d Franc) Times(multiplier int64) Franc {
    return NewFranc(d.amount * multiplier)
}

func (f Franc) Equals(mi MoneyInterface) bool {
    return f.Money.Equals(mi) &&
        reflect.TypeOf(f).Name() == reflect.TypeOf(mi).Name()
}
