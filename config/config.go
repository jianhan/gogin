package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"sync"
)

type Envs struct {
	Addr              string `env:"ADDR" envDefault:":8888"`
	ReadTimeout       int    `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout      int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	ReadHeaderTimeout int    `env:"READ_HEADER_TIMEOUT" envDefault:"10"`
	IdleTimeout       int    `env:"IDLE_TIMEOUT" envDefault:"0"`
	MaxHeaderBytes    int    `env:"MAX_HEADER_BYTES" envDefault:"0"`
	GoogleMapAPIKey   string `env:"GOOGLE_MAP_API_KEY,required"`
	ZomatoAPIKey      string `env:"ZOMATO_API_KEY,required"`
	ZomatoAPIUrl      string `env:"ZOMATO_API_URL,required"`
}

var (
	envInstance     Envs
	envInstanceOnce sync.Once
)

func GetEnvs() Envs {
	envInstanceOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		envInstance = Envs{}
		if err := env.Parse(&envInstance); err != nil {
			panic(err)
		}
	})

	return envInstance
}
