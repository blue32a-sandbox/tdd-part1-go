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

func (m Money) Equals(money Money) bool {
    return m.amount == money.amount &&
        m.currency == money.currency
}

func (m Money) Times(multiplier int64) Money {
    return Money{m.amount * multiplier, m.currency}
}

func NewDollar(amount int64) Money {
    return Money{amount, "USD"}
}

func NewFranc(amount int64) Money {
    return Money{amount, "CHF"}
}
