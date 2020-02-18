package urlRw

import (
	"encoding/json"
	"fmt"
	"github.com/saeidraei/go-crawler-clean/domain"
	_ "github.com/spf13/viper"

	"github.com/go-redis/redis/v7"
	"github.com/saeidraei/go-crawler-clean/uc"
)

type rw struct {
	client *redis.Client
}

func New() uc.QueueRW {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rw{
		client: client,
	}
}
func (rw rw) Enqueue(key string, value domain.Url) error {
	b, _ := json.Marshal(value)
	err := rw.client.LPush(key, b).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rw rw) Dequeue(key string) (*domain.Url, error) {
	val, err := rw.client.RPop(key).Result()
	if err != nil {
		return nil,err
	}
	var url domain.Url
	err = json.Unmarshal([]byte(val), &url)
	if err != nil {
		fmt.Println("error:", err)
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (rw rw) All(key string) ([]*domain.Url, error) {
	values, err := rw.client.LRange(key, 0, -1).Result()
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	var urls  []*domain.Url
	for _, val := range values {
		var url domain.Url
		err = json.Unmarshal([]byte(val), &url)
		urls = append(urls, &url)
		if err != nil {
			fmt.Println("error:", err)
		}
	}

	return urls, nil
}
