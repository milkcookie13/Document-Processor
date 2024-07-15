package store

// Конфиг для подключения к базе данных
type Config struct {
	DataBaseURL string
}

// NewConfig создает новый конфиг и возвращает его в виде указателя на структуру
func NewConfig() *Config {
	return &Config{}
}
