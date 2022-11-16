package money

type Bank struct {}

func (b Bank) reduce(source Expression, to string) Money {
    return source.Reduce(to)
}
