package main

import (
	"fmt"
)

func main() {
	s:=""
	fmt.Print("Enter moves:")
	fmt.Scan(&s)
    h,v:=0,0
	for _,m :=range s{
		switch m{
		case 'L':
			h-=1
		case 'R':
			h+=1
		case 'U':
			v+=1
		case 'D':
			v-=1
		}
	}
	if (h==0 && v==0){fmt.Println("true")}else{fmt.Println("false")}

}
