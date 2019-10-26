package algo

import (
	"fmt"
	"testing"
)

func TestDoubleRandom(t *testing.T) {
	var count,account int64
	count=3 //二个包
	account=100 //1块钱
	for i:=int64(0);i<count;i++ {
		money:=DoubleRandom(count-i,account)
		fmt.Println(money)
		account=account-money
	}
}