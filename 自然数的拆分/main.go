package main

import "fmt"

var (
	output[10]int
	input int
)

func DFS(sum,a,next int){
	if sum>input{
		return
	}
	if sum==input{
		for i:=1;i<=a-2;i++{
			fmt.Print(output[i],"+")
		}
		fmt.Print(output[a-1],"\n")
		return
	}
	for i:=next;i<=input-1;i++{
		output[a]=i
		DFS(sum+i,a+1,i)
		output[a]=0
	}
}
func main(){
	fmt.Scanf("%d",&input)
	DFS(0,1,1)
	return
}
