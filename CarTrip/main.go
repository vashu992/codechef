package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var X int
        fmt.Scan(&X)

        
        cost := 0

        if X < 300 {
            cost = 300
        } else {
            cost = X
        }

        cost *= 10 

        fmt.Println(cost)
    }
}
