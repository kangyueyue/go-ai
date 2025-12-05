package mq

import (
	"fmt"
	mylog "github.com/kangyueyue/go-ai/common/logger"
	"github.com/kangyueyue/go-ai/config"
	"github.com/streadway/amqp"
)

// coon rabbitmq connection
var conn *amqp.Connection

// initConn 初始化连接
func initConn() {
	c := config.GetConfig()
	mqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		c.RabbitMq.User,
		c.RabbitMq.Password,
		c.RabbitMq.Host,
		c.RabbitMq.Port,
		c.RabbitMq.VHost,
	)
	mylog.Log.Infof("mqUrl: %s", mqUrl)
	var err error
	conn, err = amqp.Dial(mqUrl)
	if err != nil {
		mylog.Log.Errorf("Dial mq error: %v", err)
		return
	}
}

// RabbitMq rabbitmq
type RabbitMq struct {
	coon     *amqp.Connection
	channel  *amqp.Channel
	Exchange string // 交换机
	Key      string
}

// NewRabbitMq 创建一个rabbitmq
func NewRabbitMq(exchange, key string) *RabbitMq {
	return &RabbitMq{
		Exchange: exchange,
		Key:      key,
	}
}

func (r *RabbitMq) Destroy() {
	_ = r.channel.Close()
	_ = r.coon.Close()
}

// NewWorkRabbitMq 创建一个工作队列
func NewWorkRabbitMq(queue string) *RabbitMq {
	rabbitmq := NewRabbitMq("", queue)

	// get connection
	if conn == nil {
		initConn()
	}

	// 检查连接是否成功
	if conn == nil {
		mylog.Log.Error("Failed to initialize RabbitMQ connection")
		return rabbitmq
	}

	rabbitmq.coon = conn

	// get channel
	var err error
	rabbitmq.channel, err = rabbitmq.coon.Channel()
	if err != nil {
		mylog.Log.Errorf("Get channel error: %v", err)
	}
	return rabbitmq
}

// Publish 发送消息
func (r *RabbitMq) Publish(message string) error {
	// 检查channel是否为nil
	if r.channel == nil {
		return fmt.Errorf("RabbitMQ channel is nil, cannot publish message")
	}

	_, err := r.channel.QueueDeclare(r.Key, false, false, false, false, nil)
	if err != nil {
		return err
	}
	// 调用channel
	return r.channel.Publish(r.Exchange, r.Key, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// Consume 订阅消息
func (r *RabbitMq) Consume(handle func(msg *amqp.Delivery) error) {
	// 创建队列
	q, err := r.channel.QueueDeclare(r.Key, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 接受消息
	msgList, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 处理消息
	for msg := range msgList {
		if err := handle(&msg); err != nil {
			mylog.Log.Info(err.Error())
		}
	}
}
