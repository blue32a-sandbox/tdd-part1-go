package money

type Doller struct {
    amount int64
}

func (d *Doller) Times(multiplier int64) {
    d.amount *= multiplier
}
