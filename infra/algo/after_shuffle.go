package algo

import "math/rand"

//先生成红包序列再洗牌打乱顺序  防止先大后小的概率
func AfterShuffle(count,amount int64) (result []int64) {
	//剩余金额
	balance:=amount
	//生成随机红包序列
	for i:=int64(0);i<count;i++{
		money:=SimpleRand(count-i,balance)
		balance-=money
		result=append(result,money)
	}

	//重新打散
	rand.Shuffle(len(result), func(i, j int) {
		result[i],result[j]=result[j],result[i]
	})
	return
}
