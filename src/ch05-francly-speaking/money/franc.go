package money

type Franc struct {
    amount int64
}

func NewFranc(amount int64) Franc {
    return Franc{amount: amount}
}

func (d *Franc) Times(multiplier int64) Franc {
    return Franc{amount: d.amount * multiplier}
}

func (d *Franc) Equals(franc Franc) bool {
    return d.amount == franc.amount
}
