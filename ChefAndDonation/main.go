package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)

        donation := Y - X
        fmt.Println(donation)
    }
}
