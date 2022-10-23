package main

import (
    "fmt"
    "tdd/money"
)

func main() {
    five := money.NewDollar(5)
    fmt.Printf("five: %v\n", five)

    ten := five.Times(2)
    fmt.Printf("ten: %v\n", ten)
    fmt.Printf("five: %v\n", five)

    fmt.Printf("ten == ten: %t\n", ten.Equals(money.NewDollar(10)))
}
