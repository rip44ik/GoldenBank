package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables .
// Конфигурация хранит все настройки приложения.
// Значения считываются библиотекой viper из конфигурационного файла или переменных окружения.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or enviroment variable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //jsong, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
