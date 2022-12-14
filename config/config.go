package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (conf *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	_ = godotenv.Load(filenames...)
	// exception.PanicWhenError(err)
	return &configImpl{}
}
