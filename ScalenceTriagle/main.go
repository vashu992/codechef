package main

import (
    "fmt"
)

func isScalene(a, b, c int) bool {
    return a != b && b != c && a != c
}

func main() {
    var t int
    fmt.Scan(&t)

    for i := 0; i < t; i++ {
        var a, b, c int
        fmt.Scan(&a, &b, &c)

        if isScalene(a, b, c) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
