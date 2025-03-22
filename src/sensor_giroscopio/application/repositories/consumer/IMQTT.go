package consumer

type MQTTPort interface {
	ConnectAndConsume() error
}
