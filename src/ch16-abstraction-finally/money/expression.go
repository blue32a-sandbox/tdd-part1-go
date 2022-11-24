package money

type Expression interface {
    Times(int64) Expression
    Plus(Expression) Expression
    Reduce(Bank, string) Money
}
