package mq

var (
	RMQMessage *RabbitMq
)

// InitRabbitMq 初始化rabbitmq
func InitRabbitMq() {
	RMQMessage = NewWorkRabbitMq("Message")

	// 检查RabbitMQ是否初始化成功
	if RMQMessage == nil || RMQMessage.channel == nil {
		panic("Failed to initialize RabbitMQ - check your configuration")
	}

	go RMQMessage.Consume(MqMessage) // go 异步消费消息
}

// DestroyRabbitMq 销毁rabbitmq
func DestroyRabbitMq() {
	RMQMessage.Destroy()
}
