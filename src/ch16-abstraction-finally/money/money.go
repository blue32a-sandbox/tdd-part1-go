package money

type Money struct {
    amount int64
    currency string
}

func (m Money) Amount() int64 {
    return m.amount
}

func (m Money) Currency() string {
    return m.currency
}

func (m Money) Times(multiplier int64) Expression {
    return Money{m.amount * multiplier, m.currency}
}

func (m Money) Plus(addend Expression) Expression {
    return Sum{m, addend}
}

func (m Money) Reduce(bank Bank, to string) Money {
    rate := bank.Rate(m.currency, to)
    return Money{m.amount / rate, to}
}

func NewDollar(amount int64) Money {
    return Money{amount, "USD"}
}

func NewFranc(amount int64) Money {
    return Money{amount, "CHF"}
}
