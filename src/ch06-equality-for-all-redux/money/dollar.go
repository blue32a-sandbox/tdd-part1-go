package money

type Dollar struct {
    Money
}

func NewDollar(amount int64) Dollar {
    return Dollar{Money{amount: amount}}
}

func (d Dollar) Times(multiplier int64) Dollar {
    return NewDollar(d.amount * multiplier)
}
