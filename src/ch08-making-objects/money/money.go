package money

type MoneyInterface interface {
    Amount() int64
    Equals(MoneyInterface) bool
    Times(int64) MoneyInterface
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

func NewDollar(amount int64) MoneyInterface {
    return Dollar{Money{amount: amount}}
}

func NewFranc(amount int64) MoneyInterface {
    return Franc{Money{amount: amount}}
}
