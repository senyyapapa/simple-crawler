package config

type Config struct {
	Port   string `yaml:"port"`
	DB_URL string `env:"DB_URL"`
}

func MustLoad() {

}
