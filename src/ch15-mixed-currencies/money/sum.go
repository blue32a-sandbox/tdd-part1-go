package money

type Sum struct {
    augend Expression
    addend Expression
}

func (s Sum) Plus(addend Expression) Expression {
    return nil
}

func (s Sum) Reduce(bank Bank, to string) Money {
    amount := s.augend.Reduce(bank,to).amount +
        s.addend.Reduce(bank, to).amount
    return Money{amount, to}
}
