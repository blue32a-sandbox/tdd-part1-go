package money

type MoneyInterface interface {
    Amount() int64
    Equals(MoneyInterface) bool
}

type Money struct {
    amount int64
}

func (m Money) Amount() int64 {
    return m.amount
}

func (m Money) Equals(mi MoneyInterface) bool {
    return m.amount == mi.Amount()
}
