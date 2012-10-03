package main
import(
	"fmt"
)
func fibbonacci(n int) int{
	var out int
	if n < 2 {
		if n>=0 {
			out= n
		} else {
			out= fibbonacci(n+2)-fibbonacci(n+1)
		}
	}else{
		out= fibbonacci(n-1)+ fibbonacci(n-2)
	}
	return out
}
func main(){
	for i:=0;i<40;i++{
		fmt.Println(fibbonacci(i))
	}
}
