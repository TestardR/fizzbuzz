package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/kelseyhightower/envconfig"
)

var errParseEnv = errors.New("failed to parse environment variable")

type Conf struct {
	// App config
	Port string `required:"true"`
	Env  string `required:"true" validate:"eq=debug|eq=release"`
	// Redis config
	RedisHost string `required:"true" split_words:"true"`
	RedisPort string `required:"true" split_words:"true"`
}

func New() (Conf, error) {
	var c Conf

	if err := envconfig.Process("", &c); err != nil {
		return Conf{}, fmt.Errorf("%w: %s", errParseEnv, err)
	}

	if err := validator.New().Struct(&c); err != nil {
		return Conf{}, fmt.Errorf("%w: %s", errParseEnv, err)
	}

	return c, nil
}
