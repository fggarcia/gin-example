package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"gin-example/model"
	"gin-example/util"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	_ "github.com/proullon/ramsql/driver"
	"sync"
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

	db, _ = sql.Open("ramsql", "TestLoadUserAddresses")
)

const (
	okStatus = "{status: ok}"
	albumStr = `{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99}`
	album2   = `{"id":"1","title":"The Dark Side of the Moon","artist":"Pink Floyd","price":10.99}`
)

var (
	entityPool = &sync.Pool{
		New: func() interface{} {
			return &model.Album{}
		},
	}
	album2Bytes = util.ToBytes(album2)
)

/*
curl -v --max-time 3 -X GET 'http://localhost:8080/ping'
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("context: %+v\n", c.Request.Context())
		time.Sleep(5 * time.Second)

		select {
		case <-c.Request.Context().Done():
			fmt.Printf("client gave up: %v\n", c.Request.Context().Err())
			// client gave up
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8080")
}*/

func main() {
	//db.SetMaxOpenConns(1)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/query", querySQL)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/get/:p", getOneAlbum)
	router.GET("/albums/create/:count", createAlbums)
	router.GET("/albums/delete/:count", deleteAlbums)
	router.GET("/albums/sync", syncAlbums)
	router.GET("/albums/cache", cacheSize)
	router.GET("/albums/gc", gc)
	router.GET("/albums/file/:filename", fromFiles)
	router.GET("/albums/file_fast_cache/:filename", fromFilesFastCache)
	pprof.Register(router)
	runtime.SetBlockProfileRate(1)
	router.Run("localhost:8080")
}

func syncAlbums(c *gin.Context) {
	albumPtr := entityPool.Get().(*model.Album)
	defer entityPool.Put(albumPtr)

	_ = json.Unmarshal(album2Bytes, albumPtr)
	fmt.Printf("albumPtr: %v %p\n", albumPtr, albumPtr)
	c.IndentedJSON(200, albumPtr)
}

func querySQL(c *gin.Context) {
	fmt.Printf("Calling SQL")
	batch := []string{
		`CREATE TABLE address (id BIGSERIAL PRIMARY KEY, street TEXT, street_number INT);`,
		`CREATE TABLE user_addresses (address_id INT, user_id INT);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Victor Hugo', 32);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la République', 23);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Charles Martel', 5);`,
		`INSERT INTO address (street, street_number) VALUES ('chemin du bout du monde ', 323);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la liberté', 2);`,
		`INSERT INTO address (street, street_number) VALUES ('avenue des champs', 12);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 2);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 3);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 4);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 5);`,
	}

	for _, b := range batch {
		//fmt.Printf("Calling SQL with %s\n", b)
		_, err := db.Exec(b)
		if err != nil {
			fmt.Printf("sql.Exec: Error: %s\n", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			rows, err := db.Query(`SELECT user_id FROM user_addresses;`)
			if err != nil {
				fmt.Printf("sql.Exec: Error: %s\n", err)
			}
			fmt.Printf("sql.OpenConnections() :%d\n", db.Stats().OpenConnections)
			fmt.Printf("sql.MaxOpenConnections() :%d\n", db.Stats().MaxOpenConnections)

			for rows.Next() {
				var userId string
				err := rows.Scan(&userId)
				if err == nil {
					//fmt.Printf("Rows userId: %d\n", userId)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Done\n")
	c.IndentedJSON(http.StatusOK, okStatus)
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
