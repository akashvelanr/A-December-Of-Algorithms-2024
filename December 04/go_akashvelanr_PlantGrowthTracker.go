package main
import "fmt"
func main() {
  fmt.Print("Enter no. of months:")
  var n int
  fmt.Scan(&n)
  k:=1
  sum:=1
  for i:=3;i<=n;i++{
    sum,k=sum+k,sum
    //fmt.Println(sum)
  }
  fmt.Println("Output:",sum)
}