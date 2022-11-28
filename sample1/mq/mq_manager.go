package mq

import (
	"fmt"

	amqp "github.com/streadway/amqp"
)

type MQManager struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var MQManagerInstance *MQManager

func init() {
	fmt.Println("init mq")
	MQManagerInstance = &MQManager{}
	conn, err := amqp.Dial("amqp://")
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
