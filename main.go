package main

import (
	"gin-example/model"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	//"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var (
	cache = model.NewCacheMap()
)

func main() {
	router := gin.Default()
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/create/:count", createAlbums)
	router.GET("/albums/delete/:count", deleteAlbums)
	router.GET("/albums/gc", gc)
	router.GET("/albums/file/:filename", fromFiles)
	pprof.Register(router)
	router.Run("localhost:8080")
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := cache.Get(id)
	if album == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func createAlbums(c *gin.Context) {
	count := strings.Split(c.Param("count"), ",")
	countInit, _ := strconv.Atoi(count[0])
	countEnd, _ := strconv.Atoi(count[1])
	log.Printf("count: %s", count)
	albums := make([]*album, 0, countEnd-countInit)
	for i := countInit; i < countEnd; i++ {
		album := &album{
			ID:     strconv.Itoa(i),
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		}
		albums = append(albums, album)
	}
	putIntoCache(albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func deleteAlbums(c *gin.Context) {
	count := strings.Split(c.Param("count"), ",")
	countInit, _ := strconv.Atoi(count[0])
	countEnd, _ := strconv.Atoi(count[1])
	for i := countInit; i < countEnd; i++ {
		cache.Delete(strconv.Itoa(i))
	}
	c.IndentedJSON(http.StatusOK, "{status: ok}")
}

func putIntoCache(data []*album) bool {
	for _, album := range data {
		cache.Put(album.ID, album)
	}
	return true
}

func gc(c *gin.Context) {
	log.Println("Running gc")
	runtime.GC()
	c.IndentedJSON(http.StatusOK, "{status: ok}")
}

func unmarshal(data []byte) ([]*album, error) {
	var albums []*album
	if err := json.Unmarshal(data, &albums); err != nil {
		return nil, err
	}
	return albums, nil
}
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

	albums, err := unmarshal(data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	putIntoCache(albums)

	c.IndentedJSON(http.StatusOK, "{status: ok}")
}
