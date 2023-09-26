package main

import (
    "fmt"
    "strings"
)

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var sets, tests, editorials int
        fmt.Scan(&sets, &tests, &editorials)

        category := "SETTER"
        if tests > sets && tests > editorials {
            category = "TESTER"
        } else if editorials > sets && editorials > tests {
            category = "EDITORIALIST"
        }

        fmt.Println(strings.ToUpper(category))
    }
}
