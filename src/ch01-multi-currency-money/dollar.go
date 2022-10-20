package money

type Dollar struct {
    amount int64
}

func (d *Dollar) Times(multiplier int64) {
    d.amount *= multiplier
}
