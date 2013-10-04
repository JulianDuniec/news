package store

import(
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	server = ":6379"
	newskey = "news"
	pool = &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
)

func All(from, to int) []News {
	conn := pool.Get()
	defer conn.Close()
	return newsListFromRedisValues(
		redis.Values(
			conn.Do(
				"LRANGE", 
				newskey, 
				from, 
				to)))
}

func Add(news News) {
	conn := pool.Get()
	defer conn.Close()	
	conn.Do(
		"LPUSH", 
		newskey, 
		news.serialize())
}

func newsListFromRedisValues(values []interface{}, err error) []News {
	news := []News{}
	imax := len(values)
	for i := 0; i < imax; i++ {
		buffer, _ := redis.Bytes(values[i], err)
		news = append(news, newsFromBytes(buffer))
	}	
	return news
}