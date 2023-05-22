package config

import (
	"errors"
)

type Config struct {
	PortHTTP    string
	MetricsHTTP string
}

var (
	ErrReadConfig = errors.New("error read config")
	ErrEmptyField = errors.New("error empty field")
)

func InitConfig(path string) (*Config, error) {
	//viper.AddConfigPath("config")
	//viper.SetConfigName("config")
	//viper.SetConfigFile(path)
	//fmt.Println(path)
	//fmt.Println(os.Getwd())
	//viper.AutomaticEnv()
	//if err := viper.ReadInConfig(); err != nil {
	//	return nil, ErrReadConfig
	//}
	//portHTTP := viper.GetString("app.http.port")
	//portMetrics := viper.GetString("app.metric.port")
	//if portHTTP == "" {
	//	portHTTP = ":8080"
	//}
	//if portMetrics == "" {
	//	portMetrics = ":8081"
	//}

	return &Config{PortHTTP: ":8080", MetricsHTTP: ":8081"}, nil
}
