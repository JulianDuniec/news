package main

import(
	//"github.com/julianduniec/news/server"
	"github.com/julianduniec/news/store"
	"github.com/garyburd/redigo/redis"
	"fmt" 
)


func main() {
	a := store.News{"a", "b", "c", "d"}
	conn, _ := redis.Dial("tcp", ":6379")
	defer conn.Close()
    conn.Do("SET", "key", a)

   	c, _ := conn.Do("GET", "key")
  	fmt.Println(c.(store.News))
   //server.Start()
}