package money

type Sum struct {
    augend Expression
    addend Expression
}

func (s Sum) Times(multiplier int64) Expression {
    return Sum{s.augend.Times(multiplier), s.addend.Times(multiplier)}
}

func (s Sum) Plus(addend Expression) Expression {
    return Sum{s, addend}
}

func (s Sum) Reduce(bank Bank, to string) Money {
    amount := s.augend.Reduce(bank, to).amount +
        s.addend.Reduce(bank, to).amount
    return Money{amount, to}
}
