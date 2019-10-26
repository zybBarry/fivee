package algo

import (
	"math/rand"
	"time"
)

//此算法减少了有特别大的金额出现的概率
func DoubleRandom(count,amount int64) int64 {
	if count==1{
		return amount
	}

	//计算最大可用金额
	max:=amount-min*count
	//计算第一次随机数  计算种子金额作为基数
	rand.Seed(time.Now().UnixNano())
	seed:=rand.Int63n(count*2)+1
	money:=max/seed+min
	//二次随机 计算红包金额序列
	result:=rand.Int63n(money)+min
	return result
}