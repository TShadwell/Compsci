package main
import(
	"fmt"
)
func main(){
	var total uint64
	var stack = []uint64{
		0,
		1}
	for {
		oldVal:=stack[1]
		stack[1]=stack[1]+stack[0]
		stack[0]=oldVal
		//Is beyond limit?
		if stack[1]>4000000{
			break
		//Is even?
		} else if stack[1] % 2 ==0{
			total += stack[1]
		}
	}
	fmt.Println(total)
}
