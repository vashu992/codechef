package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var A, B int
        fmt.Scan(&A, &B)

        remainder := A % B
        fmt.Println(remainder)
    }
}
