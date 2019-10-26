package algo

import (
	"fmt"
	"testing"
)

func TestDoubleAvage(t *testing.T) {
	var count,account int64
	count=3 //三个包
	account=100 //1块钱
	for i:=int64(0);i<count;i++ {
		money:=DoubleAvage(count-i,account)
		fmt.Println(money)
		account=account-money
	}
}
