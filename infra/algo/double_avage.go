package algo

import (
	"math/rand"
	"time"
)

//二倍均值算法
func DoubleAvage(count,amount int64) int64{
	if count==1{
		return  amount
	}
	//计算最大可用金额
	max:=amount-count*min
	//计算最大金额的均值
	avg:=max/count
	//计算二倍均值 +最小金额防止出现零值情况
	doubleAvg:=avg*2+min
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(doubleAvg)+min
}
