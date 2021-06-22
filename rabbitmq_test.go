package lhktools

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)
//work模式 不会重复消费
func TestV1(t *testing.T) {
	go func() {
		rabbitmq := NewRabbitMqSimple("imoocSimple")
		for d := range rabbitmq.ConsumeSimple() {
			fmt.Printf("第一个消费者 %s\n", d.Body)
		}
	}()

	go func() {
		rabbitmq := NewRabbitMqSimple("imoocSimple")
		for d := range rabbitmq.ConsumeSimple() {
			fmt.Printf("第二个消费者 %s\n", d.Body)
		}
	}()
	rabbitmq1 := NewRabbitMqSimple("imoocSimple")
	rabbitmq2 := NewRabbitMqSimple("imoocSimple")
	for i := 0; i <= 100; i++ {
		rabbitmq1.PublishSimple("生产1线-" + strconv.Itoa(i))
		rabbitmq2.PublishSimple("生产2线-" + strconv.Itoa(i))
		//time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}

func TestRouter(t *testing.T) {
	go func() {
		rabbitmq := NewRabbitMqRouter("EXCHANGE1", "success")
		for d := range rabbitmq.RecieveRouting() {
			fmt.Printf("第一个消费者 %s\n", d.Body)
		}
	}()

	rabbitmq1 := NewRabbitMqRouter("EXCHANGE1", "success")
	rabbitmq2 := NewRabbitMqRouter("EXCHANGE1", "success")
	for i := 0; i <= 100; i++ {
		rabbitmq1.PublishRouting("生产1线-success" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("生产2线-success" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}

func TestNewTopic(t *testing.T) {
	go func() {
		rabbitmq := NewRabbitMqTopic("topicExChange", "com.liu.*")
		for d := range rabbitmq.RecieveTopic() {
			fmt.Printf("第一个消费者 %s\n", d.Body)
		}
	}()

	rabbitmq1 := NewRabbitMqTopic("topicExChange", "com.liu.wang")
	rabbitmq2 := NewRabbitMqTopic("topicExChange", "com.kai.xx")
	for i := 0; i <= 100; i++ {
		rabbitmq1.PublishTopic("生产1线" + strconv.Itoa(i))
		rabbitmq2.PublishTopic("生产2线" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
