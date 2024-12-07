package main

import "fmt"

func main() {
	r,b:=0,0
	fmt.Print("Enter R:")
	fmt.Scan(&r)
	fmt.Print("Enter B:")
	fmt.Scan(&b)
	str:=""
	var m, n string

	if (r-b>1||r-b<(-1)){
		fmt.Println("Not possible")
	}else{
		if r>=b{
			m ="R"
			n ="B"
		}else{
			m,n ="B","R"
		}
		for i:=0;i<((r+b)/2)+1;i++{
			str+=m+n
		}
		fmt.Println("Output:",str[0:r+b])
	}
}
