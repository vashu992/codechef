package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)

        initialCost := X - Y
        increasedPrice := X + (X * 10 / 100)
        newProfit := increasedPrice - initialCost

        fmt.Println(newProfit)
    }
}
