// We have populated the solutions for the 10 easiest problems for your support. 
// Click on the SUBMIT button to make a submission to this problem.

package main
import "fmt"

func main(){
	
	var t int 
	fmt.Scan(&t);
	
	//println(t)
	
	for t>0 {
	    var num int
	    fmt.Scan(&num)
	    
	    var sum int=0
	    
	    for num>0 {
	        sum=sum+num%10
	        num=num/10
	    }
	    
	    fmt.Println(sum)
	    
	    t--
	}
	
}

