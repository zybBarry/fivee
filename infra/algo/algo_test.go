package algo

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleRand2(t *testing.T) {
	var count,account int64
	count=2 //二个包
	account=100 //1块钱
	total:=account
	sum:=int64(0)
	for i:=int64(0);i<2;i++ {
		money:=SimpleRand(count-i,account)
		fmt.Println(money)
		account=account-money
		sum+=money
	}
	Convey("简单随机算法",t, func() {
		So(sum,ShouldEqual,total+1)
	})
}

func AlgoTest(message string,t *testing.T,fn func(count,account int64)int64)  {
	var count,account int64
	count=2 //二个包
	account=100 //1块钱
	total:=account
	sum:=int64(0)
	for i:=int64(0);i<2;i++ {
		money:=fn(count-i,account)
		fmt.Println(money)
		account=account-money
		sum+=money
	}
	Convey(message,t, func() {
		So(sum,ShouldEqual,total)
	})
}

func TestAfterShuffle2(t *testing.T) {
	AlgoTest("先洗牌算法测试",t,BeforeShuffle)
}