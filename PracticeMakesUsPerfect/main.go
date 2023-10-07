package main
import "fmt"

func main(){
	var p1 , p2 , p3 , p4 int
	fmt.Scan(&p1 , &p2 , &p3 , &p4)
	
	weeks := 0
	
	if p1 >= 10 {
	  weeks++
	}
	if p2 >= 10 {
	  weeks++
	}
	if p3 >= 10 {
	  weeks++
	}
	if p4 >= 10 {
	  weeks++
	}
	
	fmt.Println(weeks)
}
