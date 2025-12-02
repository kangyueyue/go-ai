package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// MainConfig is the main configuration
type MainConfig struct {
	Port    int    `toml:"port"`
	AppName string `toml:"appName"`
	Host    string `toml:"host"`
	Level   string `toml:"level"`
}

// EmailConfig is the email configuration
type EmailConfig struct {
	AuthCode string `toml:"authCode"`
	Email    string `toml:"email"`
}

// RedisConfig is the redis configuration
type RedisConfig struct {
	RedisPort     int    `toml:"port"`
	RedisHost     string `toml:"host"`
	RedisDb       int    `toml:"db"`
	RedisPassword string `toml:"password"`
}

// MysqlConfig is the mysql configuration
type MysqlConfig struct {
	MysqlHost     string `toml:"host"`
	MysqlPort     int    `toml:"port"`
	MysqlDb       string `toml:"db"`
	MysqlUser     string `toml:"user"`
	MysqlPassword string `toml:"password"`
	MysqlCharset  string `toml:"charset"`
}

// JwtConfig is the jwt configuration
type JwtConfig struct {
	ExpireDuration int    `toml:"expireDuration"`
	Issuer         string `toml:"issuer"`
	Subject        string `toml:"subject"`
	Secret         string `toml:"secret"`
}

// RabbitMq is the rabbitmq configuration
type RabbitMq struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	VHost    string `toml:"vhost"`
}

// Config is the configuration
type Config struct {
	MainConfig  `toml:"main"`
	EmailConfig `toml:"email"`
	RedisConfig `toml:"redis"`
	MysqlConfig `toml:"mysql"`
	JwtConfig   `toml:"jwt"`
	RabbitMq    `toml:"rabbitmq"`
}

// RedisKeyConfig is the redis key configuration
type RedisKeyConfig struct {
	CaptchaPrefix string
}

// DefaultRedisKeyConfig is the default redis key configuration
var DefaultRedisKeyConfig = RedisKeyConfig{
	CaptchaPrefix: "captcha:%s",
}

var config *Config

// InitConfig is the initialization function
func InitConfig() error {
	// 设置文件路径，相对于main.go的路径
	if _, err := toml.DecodeFile("config/config.toml", config); err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}
	return nil
}

// GetConfig is the getter function
func GetConfig() *Config {
	if config == nil {
		config = new(Config)
		// init
		_ = InitConfig()
	}
	return config
}
