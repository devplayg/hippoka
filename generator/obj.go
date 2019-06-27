package generator

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

type Event struct {
	No      int64  `json:"i"`
	Data    string `json:"data"`
	encoded string
}

func NewEvent(no int64, data string) *Event {
	e := Event{
		No:   no,
		Data: data,
	}
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	e.encoded = string(b)

	return &e
}

//
//func NewFakeEvent(no int64, data string) *Message {
//	b, err := json.Marshal(&Event{No: no, Data: data})
//	if err != nil {
//		panic(err)
//	}
//
//	return &Message{string(b)}
//}

type Message struct {
	data string
}

func (m Message) Encode() ([]byte, error) {
	return []byte(m.data), nil
}

func (m Message) Length() int {
	return len(m.data)
}

func NewMessage(s string) *Message {
	return &Message{data: s}
}

func NewFakeMessage(topic string, partition int32, value sarama.Encoder) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Value:     value,
	}
}
