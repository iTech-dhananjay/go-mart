package config

type Config struct {
	MongoDBURI   string
	DatabaseName string
}

func NewConfig() *Config {
	return &Config{
		// Set your MongoDB URI and database name here
		MongoDBURI:   "mongodb://localhost:27017",
		DatabaseName: "yourdatabase",
	}
}
