package model

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
	Decode([]byte) interface{}
}

type KeyExtractor func(any) (string, error)

type KVSnapshot interface {
	Name() string
	Get(k string) (any, error)
	Set(data []any, f KeyExtractor) (bool, []error)
}
