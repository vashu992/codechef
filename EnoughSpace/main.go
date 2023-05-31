package main

import "fmt"

func main() {
    var T, N, X, Y int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&N, &X, &Y)

        totalSize := X + (Y * 2)

        if totalSize <= N {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
