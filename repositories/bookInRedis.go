package repositories

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"log"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
func Set(book *Book, client *redis.Client) error {
	str, err := json.Marshal(book)
	if err != nil {
		return err
	}
	log.Println(str)
	return client.Set(book.Id, str, 0).Err()

}
func Get(id string, client *redis.Client) (book *Book, err error) {
	str, err := client.Get(id).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(str), book)
	return
}
