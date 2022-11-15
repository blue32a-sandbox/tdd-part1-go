package money

type Bank struct {
    rates map[Pair]int64
}

func NewBank() Bank {
    return Bank{make(map[Pair]int64)}
}

func (b Bank) Reduce(source Expression, to string) Money {
    return source.reduce(b, to)
}

func (b Bank) AddRate(from string, to string, rate int64) {
    b.rates[Pair{from, to}] = rate
}

func (b Bank) Rate(from string, to string) int64 {
    if from == to {
        return 1
    }
    return b.rates[Pair{from, to}]
}
