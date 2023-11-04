package main
import "fmt"

func main(){
	var T , X , Y , Z int
	fmt.Scan(&T)
	
	for i := 0; i < T; i++ {
	  fmt.Scan(&X , &Y , &Z)
	  maxMangoes := (Z - Y) / X
	  fmt.Println(maxMangoes)
	}
}
