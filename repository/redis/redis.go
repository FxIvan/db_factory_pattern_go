package redis

import (
	"context"
	"fmt"

	"github.com/fxivan/db_go_abstract_factory/configuration"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	connRedis *redis.Client
}

func New(config *configuration.Configuration) (*Redis, error) {
	ctx := context.Background()
	strConn := fmt.Sprintf("%s:%d", config.Host, config.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     strConn,
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		connRedis: redisClient,
	}, nil
}

func (r *Redis) Find(id int) string {
	return "Find"
}

func (r *Redis) FindKey(key string) string {
	val, err := r.connRedis.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}

	return val
}

func (r *Redis) Save(data string) error {
	err := r.connRedis.Set(context.Background(), "comida", data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
