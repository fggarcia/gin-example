package singleflight

import (
	"github.com/goccy/go-json"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func toJsonEncode(entity *Album, bytes []byte) error {
	return json.Unmarshal(bytes, entity)
}
