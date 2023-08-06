package config

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		User:     "postgres",
		Password: "postgres",
		DBName:   "todo_db",
		SSLMode:  "disable",
	}
}

const (
	Port = ":8080"
)
