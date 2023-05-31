package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var A, B, C int
        fmt.Scan(&A, &B, &C)

        sum := A + B
        if sum == C {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
