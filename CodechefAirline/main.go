package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)
    
    for t := 0; t < T; t++ {
        var X, Y, Z int
        fmt.Scan(&X, &Y, &Z)
        
       // capacity := X
        passengers := Y
        cost := Z
        
        availableSeats := X * 10  
        if passengers <= availableSeats {
    
            totalEarnings := passengers * cost
            fmt.Println(totalEarnings)
        } else {
            
            totalEarnings := availableSeats * cost
            fmt.Println(totalEarnings)
        }
    }
}
