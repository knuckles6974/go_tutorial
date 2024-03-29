package main

import (
	"fmt"
	"runtime"
)

type count struct{
	num int
}


func (c *count) increment(){
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}
func main() {
	//실행흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장중요
	//동기화 사요하지 않은 경우 예제
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num :0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++{
		go func(){
			c.increment()
			done <- true
			runtime.Gosched() //cpu 양보

		}()
	}

	for i := 1; i <= 10000; i++{
		<-done
	}

	c.result()
}