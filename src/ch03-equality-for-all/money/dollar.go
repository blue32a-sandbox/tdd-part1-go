package money

type Dollar struct {
    Amount int64
}

func (d *Dollar) Times(multiplier int64) Dollar {
    return Dollar{Amount: d.Amount * multiplier}
}

func (d *Dollar) Equals(dollar Dollar) bool {
    return d.Amount == dollar.Amount
}
