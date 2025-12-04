package mq

var (
	RMQMessage *RabbitMq
)

// InitRabbitMq 初始化rabbitmq
func InitRabbitMq() {
	RMQMessage = NewWorkRabbitMq("Message")
	go RMQMessage.Consume(MqMessage) // go 异步消费消息
}

// DestroyRabbitMq 销毁rabbitmq
func DestroyRabbitMq() {
	RMQMessage.Destroy()
}
