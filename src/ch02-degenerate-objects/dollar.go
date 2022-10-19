package money

type Doller struct {
    amount int64
}

func (d *Doller) times(multiplier int64) Doller {
    return Doller{amount: d.amount * multiplier}
}
