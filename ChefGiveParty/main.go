package main

import (
    "fmt"
)

func main() {
    var t int
    fmt.Scan(&t) 

    for i := 0; i < t; i++ {
        var n, x, k int
        fmt.Scan(&n, &x, &k) 

        if (n*x <= k) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
