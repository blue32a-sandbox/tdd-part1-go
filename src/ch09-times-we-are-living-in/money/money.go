package money

type MoneyInterface interface {
    Amount() int64
    Currency() string
    Equals(MoneyInterface) bool
    Times(int64) MoneyInterface
}

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

func (m Money) Equals(mi MoneyInterface) bool {
    return m.amount == mi.Amount()
}

func NewMoney(amount int64, currency string) Money {
    return Money{amount, currency}
}

func NewDollar(amount int64) MoneyInterface {
    return Dollar{NewMoney(amount, "USD")}
}

func NewFranc(amount int64) MoneyInterface {
    return Franc{NewMoney(amount, "CHF")}
}
