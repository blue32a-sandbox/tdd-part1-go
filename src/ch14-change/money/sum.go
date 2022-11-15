package money

type Sum struct {
    augend Money
    addend Money
}

func (s Sum) reduce(bank Bank, to string) Money {
    amount := s.augend.amount + s.addend.amount
    return Money{amount, to}
}
