package encoder

import (
	"gin-example/model"
	"github.com/goccy/go-json"
)

type Encoder[T any] interface {
	Encode(T) ([]byte, error)
	Decode([]byte) (*T, error)
	DecodeTo([]byte, *T) error
}

type JsonEncoder[T any] struct{}

func (e JsonEncoder[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (e JsonEncoder[T]) Decode(data []byte) (*T, error) {
	var instance T
	err := json.Unmarshal(data, &instance)
	return &instance, err
}

func (e JsonEncoder[T]) DecodeTo(data []byte, instance *T) error {
	err := json.Unmarshal(data, instance)
	return err
}

func oldEncoder(instance interface{}) ([]byte, error) {
	return json.Marshal(instance)
}

func oldDecoder(data []byte) (*model.Album, error) {
	var album model.Album
	err := json.Unmarshal(data, &album)
	return &album, err
}
