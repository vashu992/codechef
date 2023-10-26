package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for t := 0; t < T; t++ {
        var X int
        fmt.Scan(&X)

        if X <= 15 {
            fmt.Println("Yes")
        } else {
            fmt.Println("No")
        }
    }
}
