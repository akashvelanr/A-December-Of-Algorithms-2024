package main

import "fmt"

func main() {
	n,k:=0,0
	fmt.Print("Enter n:")
	fmt.Scan(&n)
	fmt.Print("Enter k:")
	fmt.Scan(&k)
	s:=0

	for i:=1;i<=n;i++{
		s=(s+k)%i;
	}
	fmt.Println("Survivor:",s+1)
}
