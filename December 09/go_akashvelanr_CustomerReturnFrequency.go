package main

import (
	"fmt"
)

func main() {
	n:=0
	fmt.Print("Enter no. of returns:")
	fmt.Scan(&n)
	r := make([]int, n)
	fmt.Print("Enter all returns:")
	for i:=0;i<n;i++{fmt.Scan(&r[i])}
	s:=0

	for i:=0;i<n;i++{
		if r[i]==1{
			s+=1
		}
	}
	fmt.Println("Output:",s)
}
