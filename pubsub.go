package go_redis_interface

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

// JsonPub
func JsonPub(redisAddr string, channel string, payload interface{}) {

	var ctx = context.Background()

	var rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})

	// payload(map, structure ...) to JSON
	j, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	err = rdb.Publish(ctx, channel, j).Err()
	if err != nil {
		panic(err)
	}
}

// Sub
func Sub(goChannel chan SubMsg, redisAddr string, channel ...string) {
	var ctx = context.Background()

	var rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})

	sub := rdb.Subscribe(ctx, channel...)

	for {

		channel := sub.Channel()

		// Consume messages.
		for msg := range channel {
			goChannel <- SubMsg{
				channel: msg.Channel,
				payload: msg.Payload,
			}
		}
		log.Println("end for")
	}
}
