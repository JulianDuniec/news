package store

import(
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	/*
		Redis server
	*/
	server 		= ":6379"
	/*
		Key of the news-list
	*/
	newskey 	= "news"

	newskey_maxlen = 20
	/*
		Connection-pool for redis
	*/
	pool 		= &redis.Pool{
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

/*
	Returns all news within the 
	range, from and to
*/
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

func exists(news News) bool {
	conn := pool.Get()
	defer conn.Close()	

	exists, _ := redis.Bool(conn.Do("EXISTS", news.getId()))
	return exists
}

/*
	Adds a news-item to the list of news
*/
func Add(news News) {
	if exists(news) == true {
		return
	}
	
	conn := pool.Get()
	defer conn.Close()	

	conn.Do(
		"LPUSH", 
		newskey, 
		news.serialize())
	conn.Do(
		"SET", 
		news.getId(), 
		news.serialize())
}

/*
	Creates a []News from an []interface, returned
	from redis.Values
*/
func newsListFromRedisValues(values []interface{}, err error) []News {
	news := []News{}
	imax := len(values)
	for i := 0; i < imax; i++ {
		buffer, _ := redis.Bytes(values[i], err)
		news = append(news, newsFromBytes(buffer))
	}	
	return news
}