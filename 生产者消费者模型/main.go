package main

import "fmt"
//创建生产者，只能接受数据并传入channel中
func Producer(ch1 chan<- int){
	for i:=1;i<=9;i++{
		fmt.Println("生产者为：",i)
		ch1<-i
	}
	close(ch1)
}
//创建消费者，只能输出数据
func Consumer(ch1 <-chan int){
	for v:=range ch1{
		fmt.Println("消费者中：",v)
	}
}
func main(){
	ch:=make(chan int,10)
	go Producer(ch)
	Consumer(ch)
}
