package model

type AlbumION struct {
	ID     string `ion:"id,omitempty"`
	Title  string `ion:"title,omitempty"`
	Artist string `ion:"artist,omitempty"`
	Price  float64 `ion:"price,omitempty"`
}
