package main
import "fmt"

func main(){
	var T , S int
	fmt.Scanf("%d", &T)
	
	for i := 0; i <T; i++ {
	  fmt.Scanf("%d", &S)
	  
	  words := 1000 * 24 * S
	  fmt.Println(words)
	}
}
