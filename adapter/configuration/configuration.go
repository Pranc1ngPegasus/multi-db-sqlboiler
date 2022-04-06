package configuration

import (
	"io"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var globalConfig Config

type Config struct {
	DB1 struct {
		Database       string        `envconfig:"DATABASE1_NAME"`
		Username       string        `envconfig:"DATABASE1_USER"`
		Password       string        `envconfig:"DATABASE1_PASS"`
		Hostname       string        `envconfig:"DATABASE1_HOST"`
		Port           int           `envconfig:"DATABASE1_PORT" default:"5432"`
		SSLMode        string        `envconfig:"DATABASE1_SSL_MODE"`
		Debug          bool          `envconfig:"DATABASE1_DEBUG" default:"true"`
		ConnectTimeout int           `envconfig:"DATABASE1_CONNECTION_TIMEOUT" default:"9"`
		MaxOpen        int           `envconfig:"DATABASE1_CONNECTION_MAX_OPEN" default:"0"`
		MaxIdle        int           `envconfig:"DATABASE1_CONNECTION_MAX_IDLE" default:"2"`
		MaxLifetime    time.Duration `envconfig:"DATABASE1_CONNECTION_MAX_LIFETIME"`
	}

	DB2 struct {
		Database       string        `envconfig:"DATABASE2_NAME"`
		Username       string        `envconfig:"DATABASE2_USER"`
		Password       string        `envconfig:"DATABASE2_PASS"`
		Hostname       string        `envconfig:"DATABASE2_HOST"`
		Port           int           `envconfig:"DATABASE2_PORT" default:"5432"`
		SSLMode        string        `envconfig:"DATABASE2_SSL_MODE"`
		Debug          bool          `envconfig:"DATABASE2_DEBUG" default:"true"`
		ConnectTimeout int           `envconfig:"DATABASE2_CONNECTION_TIMEOUT" default:"9"`
		MaxOpen        int           `envconfig:"DATABASE2_CONNECTION_MAX_OPEN" default:"0"`
		MaxIdle        int           `envconfig:"DATABASE2_CONNECTION_MAX_IDLE" default:"2"`
		MaxLifetime    time.Duration `envconfig:"DATABASE2_CONNECTION_MAX_LIFETIME"`
	}
}

// Usage は各設定値の利用方法を出力する
func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

// Load は設定値を読み込む
func Load() {
	envconfig.MustProcess("", &globalConfig)
}

// Get は設定値を取得する
func Get() Config {
	return globalConfig
}
