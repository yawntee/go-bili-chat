package go_bili_chat

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	"github.com/bytedance/sonic"
)

type jsonCoder interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

type DefaultJson struct {
}

func (d *DefaultJson) Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func (d *DefaultJson) Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func SetJsonCoder(j jsonCoder) {
	serializer.JsonCoder = j
}
