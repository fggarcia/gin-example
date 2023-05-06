package main

import (
	"bytes"
	"fmt"
	"gin-example/model"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	//"encoding/json"
	//"github.com/go-json-experiment/json"
	//json "github.com/bytedance/sonic"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// albums slice to seed record album data.
var (
	cache             = model.NewCacheMap()
	encoder           = &model.AlbumEncoder{}
	fast_cache        = model.NewFastCache("album", encoder)
	extract_album_key = func(entity any) (string, error) { return entity.(model.Album).ID, nil }
)

const (
	okStatus = "{status: ok}"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/get/:p", getOneAlbum)
	router.GET("/albums/create/:count", createAlbums)
	router.GET("/albums/delete/:count", deleteAlbums)
	router.GET("/albums/cache", cacheSize)
	router.GET("/albums/gc", gc)
	router.GET("/albums/file/:filename", fromFiles)
	router.GET("/albums/file_fast_cache/:filename", fromFilesFastCache)
	pprof.Register(router)
	router.Run("localhost:8080")
}

func getOneAlbum(c *gin.Context) {
	p := c.Param("p")
	rawEntity := cache.Get("0")
	entity := encoder.Decode(rawEntity.([]byte))
	album := entity.(*model.Album)
	fmt.Printf("raw album: %v\n", album)
	album.ID = p
	c.IndentedJSON(http.StatusOK, album)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	//album := cache.Get(id)
	album, _ := fast_cache.Get(id)
	if album == nil {
		c.Status(http.StatusNotFound)
		return
	}
	//c.IndentedJSON(http.StatusOK, album)
	c.IndentedJSON(http.StatusOK, album.(*model.Album))
}

func createAlbums(c *gin.Context) {
	count := strings.Split(c.Param("count"), ",")
	countInit, _ := strconv.Atoi(count[0])
	countEnd, _ := strconv.Atoi(count[1])
	log.Printf("count: %s", count)
	albums := make([]*model.Album, 0, countEnd-countInit)
	for i := countInit; i < countEnd; i++ {
		album := &model.Album{
			ID:     strconv.Itoa(i),
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		}
		albums = append(albums, album)
	}
	putIntoCache(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func deleteAlbums(c *gin.Context) {
	count := strings.Split(c.Param("count"), ",")
	countInit, _ := strconv.Atoi(count[0])
	countEnd, _ := strconv.Atoi(count[1])
	for i := countInit; i < countEnd; i++ {
		cache.Delete(strconv.Itoa(i))
	}
	c.IndentedJSON(http.StatusOK, okStatus)
}

func putIntoCache(data *[]*model.Album) bool {
	albumsTmp := make([][]byte, 0, len(*data))
	for _, album := range *data {
		entity, err := encoder.Encode(album)
		if err != nil {
			fmt.Printf("encode error: %v", err)
		} else {
			albumsTmp = append(albumsTmp, entity)
		}
	}
	for idx, album := range albumsTmp {
		var copied = make([]byte, 0, len(album))
		copy(copied, album)
		key := strconv.Itoa(idx)
		cache.Put(key, &copied)
	}
	return true
}

func gc(c *gin.Context) {
	log.Println("Running gc")
	runtime.GC()
	c.IndentedJSON(http.StatusOK, okStatus)
}

/*
func unmarshal(data []byte) ([]*album, error) {
	var albums []*album
	if err := json.Unmarshal(data, &albums); err != nil {
		return nil, err
	}
	return albums, nil
}*/

func unmarshal(data *[]byte) (*[]*model.Album, error) {
	var albums []*model.Album
	if err := json.NewDecoder(bytes.NewReader(*data)).Decode(&albums); err != nil {
		return nil, err
	}
	return &albums, nil
}

/*
	func unmarshal(data []byte) ([]*album, error) {
		var tmp []json.RawMessage
		//if err := json.Unmarshal(data, &tmp); err != nil {
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&tmp); err != nil {
			return nil, err
		}
		albums := make([]*album, len(tmp))

		for i, raw := range tmp {
			album := &album{}
			if err := json.Unmarshal(raw, album); err != nil {
				return nil, err
			}
			albums[i] = album
		}

		return albums, nil
	}
*/
func getBytesFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func fromFiles(c *gin.Context) {
	fileName := c.Param("filename")
	data, err := getBytesFromFile(fileName + ".json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	unmarshalAndPut(c, &data)
}

func unmarshalAndPut(c *gin.Context, data *[]byte) {
	albums, err := unmarshal(data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	putIntoCache(albums)
	c.IndentedJSON(http.StatusOK, okStatus)
}

func fromFilesFastCache(c *gin.Context) {
	fileName := c.Param("filename")
	data, err := getBytesFromFile(fileName + ".json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	albumsTmp, err := unmarshal(&data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	albums := make([]any, len(*albumsTmp))
	for _, album := range *albumsTmp {
		albums = append(albums, album)
	}

	fast_cache.Set(albums, extract_album_key)

	c.IndentedJSON(http.StatusOK, okStatus)
}

type CacheSize struct {
	Size int `json:"size"`
}

func cacheSize(c *gin.Context) {
	size := cache.Size()
	c.IndentedJSON(http.StatusOK, CacheSize{Size: size})
}
