package store

import(
	
	"encoding/gob"
	"bytes"
)

type News struct {
	Url				string  		`json:"url"`
	Title			string			`json:"title"`
	Preamble 		string 			`json:"preamble"`
	Body 			string			`json:"body"`
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