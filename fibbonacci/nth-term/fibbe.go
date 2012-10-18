package main
import(
	"fmt"
)
func fibbonacci(n int) int64{
	var stack = []int64{
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
	for i:=0;i<60;i++{
		fmt.Println(fibbonacci(i))
	}
}
