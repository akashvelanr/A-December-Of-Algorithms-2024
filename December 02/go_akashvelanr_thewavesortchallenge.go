package main

import "fmt"

func main() {
	var arr= []int{10, 5, 6, 3, 2, 20, 100, 80}
    fmt.Println("Array=",arr)

	//var newarr []int
    
	for i:=0;i<len(arr)-1;i++{
		j:=i+1
        if j%2==1{
           if arr[j]>arr[i]{
			arr[j],arr[i]=arr[i],arr[j]
		   }
		}else{
			if arr[j]<arr[i]{
				arr[j],arr[i]=arr[i],arr[j]
			   }
		}

	}
	fmt.Println("Output=",arr)
}
