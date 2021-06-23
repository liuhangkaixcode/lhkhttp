package lhktools

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"math/rand"
	"time"
)

import (
	"fmt"
	"testing"
)

func TestKafka(t *testing.T) {
  go producer_test()
 // go consumer_test()

	time.Sleep(time.Second*100)
}

var(

)
func producer_test() {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key:   sarama.StringEncoder("go_test"),
	}


	for {
		//time.Sleep(time.Second)

		tt:=make(map[string]interface{})
		rand.Seed(time.Now().UnixNano())
		k:=rand.Intn(3)
		sname:=[3]string{"秒杀系统","下单系统","发货系统"}
		tt["sname"]=sname[k]
		rand.Seed(time.Now().UnixNano())
		k2:=rand.Intn(3)
		iname:=[3]string{"查询用户接口","校验接口","更新接口"}
		tt["iname"]=iname[k2]
		rand.Seed(time.Now().UnixNano())
		k3:=rand.Intn(3)
		contents:=[3]string{"===||==该接口该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功该接口发送成功发送成功==||===","该接口请求失败","该接口不知道是否成功"}
		istatus:=[3]string{"success","error","unknow"}
		tt["istatus"]=istatus[k3]
		tt["msg"]=contents[k3]
		tt["time"]=time.Now().Format("2006-01-02 15:04:05")

		bytes, _ := json.Marshal(tt)

		msg.Value = sarama.ByteEncoder(string(bytes))

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

func consumer_test() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}