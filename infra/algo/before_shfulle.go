package algo

import (
	"math/rand"
	"time"
)

//先洗牌算法 会有先大后小的现象
func BeforeShuffle(count,amount int64) int64  {
	if count==1{
		return amount
	}
	//计算最大可分配金额
	max:=amount-min*count
	//生成红包种子金额序列
	seeds:=make([]int64,0)
	//红包种子金额序列长度=3~1/2*count
	size:=count/2
	if size<3{
		size=3
	}
	for i:=int64(0);i<size;i++{
		x:=max/(i+1)
		seeds=append(seeds,x)
	}
	//从红包种子金额序列中随机选出一个作为基数
	rand.Seed(time.Now().UnixNano())
	id:=rand.Int63n(int64(len(seeds)))
	money:=rand.Int63n(seeds[id])+min
	return  money
}
