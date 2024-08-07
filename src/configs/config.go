package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
}

func LoadConfig(path string) (*Conf, error) {
	var config *Conf
	viper.SetConfigName("app_config") // nome do arquivo de configuração
	viper.SetConfigType("env")        // tipo do arquivo que será lido
	viper.AddConfigPath(path)         // caminho do arquivo que será lido
	viper.SetConfigFile(".env")       // nome do arquivo que será lido
	viper.AutomaticEnv()              // Isso permite que, caso exista uma variável de ambiente definida no sistema, sobrescreva a definida no .env
	err := viper.ReadInConfig()       // lê o arquivo .env
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config) // transforma o arquivo .env em um struct
	if err != nil {
		panic(err)
	}
	return config, nil
}
