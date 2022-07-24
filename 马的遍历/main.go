package main

import "fmt"

type Struct struct{
	x,y,s int
}

var (
	n,m,x,y,head,tail int
	dx=[9]int{0,1,2,2,1,-1,-2,-2,-1}
	dy=[9]int{0,2,1,-1,-2,-2,-1,1,2}
	Horse[440*440] Struct
	result[440][440]int
)

//将结果输出的矩阵初始化为-1
func Initial(){
	for i:=0;i<=400;i++{
		for j:=0;j<=400;j++{
			result[i][j]=-1
		}
	}
}
func BFS(){
	for {
		if head >= tail {
			break
		} else {
			head++              //头变尾了
			for i := 1; i <= 8; i++ {
				x := Horse[head].x
				y := Horse[head].y
				x = x + dx[i]
				y = y + dy[i]
				if x > 0 && y > 0 && x <= n && y <= m && result[x][y] == -1 {
					tail++
					Horse[tail].x = x
					Horse[tail].y = y
					Horse[tail].s = Horse[head].s + 1
					result[x][y] = Horse[tail].s
				}
			}
		}
	}
}
func main(){
	fmt.Scanf("%d%d%d%d",&n,&m,&x,&y)
	Initial()
	head=0
	tail=1
	result[x][y]=0
	Horse[tail].x=x
	Horse[tail].y=y
	Horse[tail].s=0
	BFS()
	for i:=1;i<=n;i++ {
		for j:=1;j<m;j++{
			fmt.Printf("%-5d",result[i][j])
		}
			fmt.Println(result[i][m])
	}
}