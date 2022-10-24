package money

type Dollar struct {
    amount int64
}

func NewDollar(amount int64) Dollar {
    return Dollar{amount: amount}
}

func (d Dollar) Times(multiplier int64) Dollar {
    return Dollar{amount: d.amount * multiplier}
}

func (d Dollar) Equals(dollar Dollar) bool {
    return d.amount == dollar.amount
}
