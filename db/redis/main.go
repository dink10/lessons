package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	status := client.Ping(ctx)
	if status.Err() != nil {
		logrus.Fatal(status.Err())
	}

	data := []struct {
		Name string
		Age  int
	}{
		{
			Name: "Anton",
			Age:  30,
		},
		{
			Name: "Kate",
			Age:  16,
		},
		{
			Name: "Vasily",
			Age:  23,
		},
	}

	for _, v := range data {
		err := client.Set(ctx, v.Name, v.Age, 0).Err()
		if err != nil {
			logrus.Error(err)
			return
		}
	}

	for _, v := range data {
		res, err := client.Get(ctx, v.Name).Int()
		if err != nil {
			logrus.Error(err)
			return
		}

		logrus.Printf("Records from Redis: %s : %d\n", v.Name, res)
	}

	d, _ := json.Marshal(data)

	err := client.Set(ctx, "test", d, 0).Err()
	if err != nil {
		logrus.Error(err)
		return
	}

	res, err := client.Get(ctx, "test").Bytes()
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Println(string(res))
}
