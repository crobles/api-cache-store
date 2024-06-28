package connectDatabase

import (
	config "api-cache-store/config"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var clientDB *redis.Client

func ConnectDB() (*redis.Client, error) {
	ctx := context.Background()
	redisAddr := config.Config("REDIS_ADDRESS")
	redisPass := config.Config("REDIS_PASSWORD")
	redisDB := config.Config("REDIS_DB")

	db, err := strconv.Atoi(redisDB)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:	redisAddr,
		Password:	redisPass,
		DB:	db,
	})
	//defer client.Close()
	status, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
		return nil, err
	}
	fmt.Println("Ping: --->", status)
	clientDB = client
	fmt.Println("Connected to Redis on ", redisAddr)
	return client, nil
}

func GetRedisClient() *redis.Client {
	return clientDB
}


