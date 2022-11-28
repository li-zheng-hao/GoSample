package mq

import (
	"fmt"

	amqp "github.com/streadway/amqp"
)

type MQManager struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (mq *MQManager) Publish(exchange, key string, msg []byte) error {
	err := mq.Channel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})
	return err
}
func (mq *MQManager) Consume(key string) {
	msgs, err := mq.Channel.Consume(key, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}
}
func (mq *MQManager) Close() {
	mq.Channel.Close()
	mq.Connection.Close()
}

var MQManagerInstance *MQManager

func init() {
	fmt.Println("init mq")
	MQManagerInstance = &MQManager{}
	conn, err := amqp.Dial("amqp://admin::5672//cloud")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
}
func initConnection() {

}

// 获取连接单例对象
func GetMQManager() *MQManager {
	if MQManagerInstance == nil {
		initConnection()
	}
	return MQManagerInstance
}
