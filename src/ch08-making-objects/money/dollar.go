package money

import "reflect"

type Dollar struct {
    Money
}

func (d Dollar) Times(multiplier int64) MoneyInterface {
    return NewDollar(d.amount * multiplier)
}

func (d Dollar) Equals(mi MoneyInterface) bool {
    return d.Money.Equals(mi) &&
        reflect.TypeOf(d).Name() == reflect.TypeOf(mi).Name()
}
