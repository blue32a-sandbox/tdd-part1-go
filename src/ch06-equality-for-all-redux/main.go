package main

import (
    "fmt"
    "tdd/money"
)

func main() {
    fiveDollar := money.NewDollar(5)
    fmt.Printf("five dollar: %v\n", fiveDollar)

    tenDollar := fiveDollar.Times(2)
    fmt.Printf("ten dollar: %v\n", tenDollar)
    fmt.Printf("five dollar: %v\n", fiveDollar)

    fmt.Printf("ten dollar == ten dollar: %t\n", tenDollar.Equals(money.NewDollar(10)))

    fiveFranc := money.NewFranc(5)
    fmt.Printf("five franc: %v\n", fiveFranc)

    tenFranc := fiveFranc.Times(2)
    fmt.Printf("ten franc: %v\n", fiveFranc)
    fmt.Printf("five franc: %v\n", tenFranc)

    fmt.Printf("ten franc == ten franc: %t\n", tenFranc.Equals(money.NewFranc(10)))
}
