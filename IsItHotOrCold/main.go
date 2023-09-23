package main
import "fmt"

func main(){
var t int 
fmt.Scan(&t)

 for i := 0; i < t; i++ {
   var c int
   fmt.Scan(&c)
   
   if c > 20 {
     fmt.Println("HOT")
   }else {
     fmt.Println("COLD")
   }
 }
}
