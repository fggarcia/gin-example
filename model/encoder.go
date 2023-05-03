package model

import (
	"encoding/json"
)

type AlbumEncoder struct {
}

func (e *AlbumEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (e *AlbumEncoder) Decode(data []byte) interface{} {
	var entity *Album
	err := json.Unmarshal(data, &entity)
	if err != nil {
		return nil
	}
	return entity
}
