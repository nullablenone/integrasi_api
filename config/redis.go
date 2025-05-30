package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // default Redis port
		Password: "",               // kosongin kalau gak ada password
		DB:       0,                // default DB
	})

	// Tes koneksi sekali pake context.Background()
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}

	log.Println("Berhasil terhubung ke Redis")
	return rdb
}
