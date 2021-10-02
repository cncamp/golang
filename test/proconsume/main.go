package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 练习2-基于 Channel 编写一个简单的单线程生产者消费者模型
func Test_06() {

	queue := make(chan int, 10)

	//生产者
	producer := time.NewTimer(time.Second)
	go func() {
		for {
			select {
			case <-producer.C:
				queue <- rand.Intn(5)
				fmt.Println("send data to proconsume")
				producer.Reset(time.Second)
			}
		}
	}()

	//消费者
	consumer := time.NewTimer(time.Second)
	for {
		select {
		case <-consumer.C:
			date := <-queue
			fmt.Println("receive data from proconsume,date:", date)
			consumer.Reset(time.Second)
		}
	}

}

func main() {
	Test_06()
}
