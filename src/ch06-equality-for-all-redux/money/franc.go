package money

type Franc struct {
    Money
}

func NewFranc(amount int64) Franc {
    return Franc{Money{amount: amount}}
}

func (d Franc) Times(multiplier int64) Franc {
    return NewFranc(d.amount * multiplier)
}
