package main

import(
	"fmt"
)

func main(){
	var total uint64
	for i:=1;i<1000;i++{
		if (i%3==0)||(i%5==0){
			total+=uint64(i)
		}
	}
	fmt.Println(total)
}
