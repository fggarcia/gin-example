package gzip

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/klauspost/compress/gzip"
	//"compress/gzip"
	"io"
	"sync"
)

const (
	albumJSON = `{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99}`
)

var (
	albumJSONBytes = []byte(albumJSON)

	smallPayloadBytes    = generateNotRandomPayloadByte(64)
	mediumPayloadBytes   = generateNotRandomPayloadByte(1024)
	largePayloadBytes    = generateNotRandomPayloadByte(1024 * 1024)
	gzipAlbum, _         = zipValue(albumJSONBytes)
	gzipSmallPayload, _  = zipValue(smallPayloadBytes)
	gzipMediumPayload, _ = zipValue(mediumPayloadBytes)
	gzipLargePayload, _  = zipValue(largePayloadBytes)

	readers sync.Pool
)

type Data struct {
	Field string `json:"field"`
}

func generateNotRandomPayloadByte(size int) []byte {
	var datas []Data
	jsonStr := ""
	for len(jsonStr) <= size {
		datas = append(datas, Data{Field: "example"})
		jsonData, _ := json.Marshal(datas)
		jsonStr = string(jsonData)
	}
	return []byte(jsonStr)
}

func zipValue(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(b); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetReader(src io.Reader) (reader *gzip.Reader) {
	if r := readers.Get(); r != nil {
		reader = r.(*gzip.Reader)
		reader.Reset(src)
	} else {
		reader, _ = gzip.NewReader(src)
	}
	return reader
}

// PutReader closes and returns a gzip.Reader to the pool
// so that it can be reused via GetReader.
func PutReader(reader *gzip.Reader) {
	reader.Close()
	readers.Put(reader)
}

func unzipPoolValue(b []byte) ([]byte, error) {
	buf := bytes.NewBuffer(b)
	gz := GetReader(buf)
	defer PutReader(gz)
	return io.ReadAll(gz)
}

func unzipValue(b []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	defer gz.Close()
	return io.ReadAll(gz)
}
