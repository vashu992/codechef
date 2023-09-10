package main
import "fmt"

func main(){
  var A, B, X, Y int
  fmt.Scan(&A, &B, &X,&Y)
  
  MESSIPOINT := (2 * A) + B 
  RONALDOPOINT := (2 * X) + Y 
  
  if MESSIPOINT > RONALDOPOINT {
    fmt.Println("MESSI")
  }else if RONALDOPOINT > MESSIPOINT{
    fmt.Println("RONALDO")
  }else {
    fmt.Println("EQUAL")
  }
}
