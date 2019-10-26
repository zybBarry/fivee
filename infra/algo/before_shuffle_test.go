package algo

import (
	"fmt"
	"testing"
)

func TestBeforeShuffle(t *testing.T) {
	var count,account int64
	count=3 //二个包
	account=100 //1块钱
	for i:=int64(0);i<count;i++ {
		money:=BeforeShuffle(count-i,account)
		fmt.Println(money)
		account=account-money
	}
}
