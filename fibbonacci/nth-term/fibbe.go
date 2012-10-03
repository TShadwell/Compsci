package main
import(
	"fmt"
)
func fibbonacci(n int) int{
	var stack = []int{
		0,
		1}
	for ;n>0;n-- {
		oldVal:=stack[1]
		stack[1]=stack[1]+stack[0]
		stack[0]=oldVal
	}
	return stack[1]
}
func main(){
	for i:=0;i<40;i++{
		fmt.Println(fibbonacci(i))
	}
}
