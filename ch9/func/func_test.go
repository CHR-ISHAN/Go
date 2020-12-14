package _func

import (
	"fmt"
	"testing"
	"time"
)

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("spend : ", time.Since(start).Seconds())
		return ret
	}
}

func SlowFun(op int) int {
	time.Sleep(time.Second * 1)
	fmt.Println("func")
	return op
}

func TestFun(t *testing.T) {
	ttFun := timeSpend(SlowFun)
	//t.Log(&ttFun)
	S := ttFun(10)
	t.Log(S)
}
