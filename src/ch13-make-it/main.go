package main

import (
    "fmt"
    "tdd/money"
)

func main() {
    fiveDollar := money.NewDollar(5)
    fmt.Printf("five dollar: %v\n", fiveDollar)
    fmt.Printf("find dollar amount: %d\n", fiveDollar.Amount())
    fmt.Printf("find dollar currency: %s\n", fiveDollar.Currency())

    tenDollar := fiveDollar.Times(2)
    fmt.Printf("ten dollar: %v\n", tenDollar)
    fmt.Printf("five dollar: %v\n", fiveDollar)

    fmt.Printf("ten dollar == ten dollar: %t\n", tenDollar.Equals(money.NewDollar(10)))

    fiveFranc := money.NewFranc(5)
    fmt.Printf("five franc: %v\n", fiveFranc)
    fmt.Printf("find franc amount: %d\n", fiveFranc.Amount())
    fmt.Printf("find franc currency: %s\n", fiveFranc.Currency())

    tenFranc := fiveFranc.Times(2)
    fmt.Printf("ten franc: %v\n", fiveFranc)
    fmt.Printf("five franc: %v\n", tenFranc)

    fmt.Printf("ten franc == ten franc: %t\n", tenFranc.Equals(money.NewFranc(10)))

    fmt.Printf("ten dollar == ten franc: %t\n", tenDollar.Equals(tenFranc))
}
