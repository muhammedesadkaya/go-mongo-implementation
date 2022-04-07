package config

type AppConfig struct {
	MongoDB MongoDBConfig `yaml:"mongodb"`
}

type MongoDBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}
