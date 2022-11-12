package money

type Expression interface {
    reduce(string) Money
}
