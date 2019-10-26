package algo

import (
	"math/rand"
	"time"
)

//最小金额为一分，避免零元红包
const min = int64(1)

//简单随机算法
//红包的数量，红包的金额
//单位金额为分
func SimpleRand(count, amount int64) int64 {
	if count == 1 {
		return amount
	}

	//本次红包最大可支配金额
	max:=amount-min*count
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max)+min
}
