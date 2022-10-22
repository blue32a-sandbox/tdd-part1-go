package main

import (
    "fmt"
    "tdd/money"
)

func main() {
    five := money.Dollar{Amount: 5}
    fmt.Printf("five: %v\n", five)

    ten := five.Times(2)
    fmt.Printf("ten: %v\n", ten)
    fmt.Printf("five: %v\n", five)
}
