package main

import "fmt"

func main() {
    fmt.Print("Enter N:")
    var N int
    fmt.Scan(&N)

	var a =make([]int,N)
	fmt.Println("Enter array elements:")
	for i:=0;i<N-1;i++{
		fmt.Scan(&a[i])
	}
    
	sum:= N*(N+1)/2
    s:=0
	for _, value := range a {
        s += value
    }

    fmt.Println(sum-s,"is absent")
}
