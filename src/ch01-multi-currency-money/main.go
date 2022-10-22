package main

import (
    "fmt"
    "tdd/money"
)

func main() {
    five := money.Dollar{Amount: 5}
    fmt.Printf("%v\n", five)

    five.Times(2)
    fmt.Printf("%v\n", five)
}
