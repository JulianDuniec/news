package store

import(
	"crypto/md5"
	"io"
	"encoding/gob"
	"bytes"
	"fmt"
	"time"
)

type News struct {
	Url				string  		`json:"url"`
	Title			string			`json:"title"`
	Preamble 		string 			`json:"preamble"`
	Body 			string			`json:"body"`
	PubDate			time.Time 		`json:"pubDate"`
}

/*
	Serializes the news into a bytearray
*/
func (n News) serialize() []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(n)
	return buffer.Bytes()
}

func (n News) getId() string {
	h := md5.New()
	io.WriteString(h, n.Url)
	return fmt.Sprintf("news_%x", h.Sum(nil))
}

/*
	Deserializes a bytearray into a news-object
*/
func newsFromBytes(b []byte) News {
	buffer := *bytes.NewBuffer(b)
	dec := gob.NewDecoder(&buffer)
	var news News
	dec.Decode(&news)
	return news
}