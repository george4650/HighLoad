package config

import "github.com/spf13/viper"

type Config struct {
	Databases Databases
	HttpPort  HttpPort
}

func Load(envPrefix string) *Config {
	// Подключение env-переменных
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	// END Подключение env-переменных

	return &Config{
		Databases: Databases{
			Oracle: DatabaseOracle{
				Host:     "",
				Port:     1521,
				Service:  "",
				User:     "",
				Password: "",
			},
		},
		HttpPort: HttpPort{
			HttpPort: 8080,
		},
	}
}
