package algo

import (
	"fmt"
	"testing"
)

func TestSimpleRand(t *testing.T) {
	var count,account int64
	count=10 //十个包
	account=100 //1块钱
	for i:=int64(0);i<10;i++ {
		money:=SimpleRand(count,account)
		fmt.Println(money)
	}
}
