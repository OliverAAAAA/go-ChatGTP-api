package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

var config *Config

type Config struct {
	ChatGpt     ChatGptConfig `json:"chatgpt" mapstructure:"chatgpt" yaml:"chatgpt"`
	MysqlConfig MysqlConfig   `json:"mysql" mapstructure:"mysql" yaml:"mysql"`
}

type ChatGptConfig struct {
	Token         string `json:"token,omitempty"  mapstructure:"token,omitempty"  yaml:"token,omitempty"`
	Ip            string `json:"ip,omitempty"  mapstructure:"ip,omitempty"  yaml:"ip,omitempty"`
	Port          string `json:"port,omitempty"  mapstructure:"port,omitempty"  yaml:"port,omitempty"`
	RequestSecret string `json:"requestSecret,omitempty"  mapstructure:"requestSecret,omitempty"  yaml:"requestSecret,omitempty"`
}

type MysqlConfig struct {
	Host     string `json:"host,omitempty"  mapstructure:"host,omitempty"  yaml:"host,omitempty"`
	Port     string `json:"port,omitempty"  mapstructure:"port,omitempty"  yaml:"port,omitempty"`
	Username string `json:"username,omitempty"  mapstructure:"username,omitempty"  yaml:"username,omitempty"`
	Password string `json:"password,omitempty"  mapstructure:"password,omitempty"  yaml:"password,omitempty"`
	Database string `json:"database,omitempty"  mapstructure:"database,omitempty"  yaml:"database,omitempty"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetMysqlConfig() *MysqlConfig {
	return &config.MysqlConfig
}

func GetIp() *string {
	ip := getEnv("ip")

	if ip != nil {
		return ip
	}
	if config == nil {
		return nil
	}
	if ip == nil {
		ip = &config.ChatGpt.Ip
	}
	return ip
}
func GetRequestSecret() *string {
	requestSecret := getEnv("requestSecret")

	if requestSecret != nil {
		return requestSecret
	}
	if config == nil {
		return nil
	}
	if requestSecret == nil {
		requestSecret = &config.ChatGpt.RequestSecret
	}
	return requestSecret
}
func GetPort() *string {
	port := getEnv("port")

	if port != nil {
		return port
	}
	if config == nil {
		return nil
	}
	if port == nil {
		port = &config.ChatGpt.Port
	}
	return port
}

func GetOpenAiApiKey() *string {
	apiKey := getEnv("api_key")

	if apiKey != nil {
		return apiKey
	}

	if config == nil {
		return nil
	}
	if apiKey == nil {
		apiKey = &config.ChatGpt.Token
	}
	return apiKey
}

func getEnv(key string) *string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = os.Getenv(strings.ToUpper(key))
	}

	if len(value) > 0 {
		return &value
	}

	if config == nil {
		return nil
	}

	return nil
}
