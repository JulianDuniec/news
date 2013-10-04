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

func (n News) serialize() []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(n)
	return network.Bytes()
}

func newsFromBytes(b []byte) News {
	network := *bytes.NewBuffer(b)
	dec := gob.NewDecoder(&network)
	var news News
	dec.Decode(&news)
	return news
}