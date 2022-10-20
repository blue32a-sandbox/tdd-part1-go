package money

type Dollar struct {
    amount int64
}

func (d *Dollar) Times(multiplier int64) Dollar {
    return Dollar{amount: d.amount * multiplier}
}

func (d *Dollar) Equals(dollar Dollar) bool {
    return d.amount == dollar.amount
}
