package money

type Expression interface {
    reduce(Bank, string) Money
}
