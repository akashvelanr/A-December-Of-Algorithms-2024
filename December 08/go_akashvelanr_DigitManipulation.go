package main

import (
	"fmt"
)

func main() {
	n:=0
	fmt.Print("Enter n:")
	fmt.Scan(&n)
	
	s:=0

	for i:=1;i<=n;i++{
		j:=i
		for j>0{

			r:=j%10
			s+=r*r
			j=j/10
		}
	}
	fmt.Println("Output:",s)
}
