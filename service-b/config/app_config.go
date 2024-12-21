package config

import "github.com/spf13/viper"

type Conf struct {
	ApiUrlZip          string `mapstructure:"API_URL_ZIP"`
	ApiUrlWeather      string `mapstructure:"API_URL_WEATHER"`
	ApiKeyWeather      string `mapstructure:"API_KEY_WEATHER"`
	WebServerPort      string `mapstructure:"WEB_SERVER_PORT"`
	ServiceName        string `mapstructure:"SERVICE_NAME"`
	CollectorUrl       string `mapstructure:"COLLECTOR_URL"`
	ServiceNameRequest string `mapstructure:"SERVICE_NAME_REQUEST"`
}

func LoadConfig() (*Conf, error) {
	viper.SetConfigName("app_b_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf Conf
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
