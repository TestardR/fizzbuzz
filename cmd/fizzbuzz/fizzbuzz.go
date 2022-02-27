package main

import (
	"errors"
	"fmt"

	"github.com/TestardR/fizz-buzz/config"
	"github.com/TestardR/fizz-buzz/internal/handler/http"
	"github.com/TestardR/fizz-buzz/internal/storage/redis"
	"github.com/TestardR/fizz-buzz/pkg/logger"
)

const appName = "fizz-buzz"

var errStoreInstance = errors.New("failed to instantiate storage")

func main() {
	log := logger.NewLogger(appName)

	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	store, err := redis.NewStorage(c.RedisHost, c.RedisPort)
	if err != nil {
		log.Fatal(fmt.Errorf("%w: %s", errStoreInstance, err))
	}

	s := http.NewServer(c.Env, log, store)
	if err := s.Run(":" + c.Port); err != nil {
		log.Fatal(err)
	}
}
