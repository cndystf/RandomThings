package main
import "fmt"

func main() {
 	//infer data type
 	var x = 1
  	x = 100
  
 	//initializing and declaring
 	y :="Golang"
  
	fmt.Println(x)
	fmt.Println(y)
	
	//aritmethic
	var st = 10
	var nd = 20
	
	fmt.Println(st+nd)
	fmt.Println(nd-st)
	fmt.Println(nd*st)
	fmt.Println(nd/st)
	
	//loop 
	i := 1
	
	for i <= 10 {
	  fmt.Println(i)
	  i = i + 1 
	}
	
	//flow control
	j := 10
	if j%2 == 0 {
	  fmt.Println("Value is even")
	} 
	else {
	  fmt.Println("Value is odd")
	}
	
	//ver 2
	k : = "Yes"
		
	if k == "Yes" {
	  fmt.Println("I say yes")
	} 
	else if k == "No" {
	  fmt.Println("I say no")
	}
	else {
	  fmt.Println("I say neither yes nor no")
	}
	
	//switch
	z := 2 
	switch z {
	  case 1 :
	    fmt.Println("z is 1")
	  case 2 :
	    fmt.Println("z is 2")
	  case 3 :
	    fmt.Println("z is 3")
	  default :
	    fmt.Println("z is neither 1, 2, or 3")
	}
	
	//array
	var a[5] int
	fmt.Println("empty array :", a)
	
	a[4] = 10
	fmt.Println(a)
	
	//slice
	s := make([]string, 3)
	fmt.Println("empty slice : ", s)
	
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	
	fmt.Println(s)
	
	s = append(s, "d")
	fmt.Println(s)
	
	//0  1 2 3 4 5
	//[a b c d e f]
	p :=s[2:5]
	fmt.Println(p)
}

