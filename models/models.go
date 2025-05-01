package helpers

type Payload struct {
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	UtcTimestamp string      `json:"utc_timestamp"`
	Data         interface{} `json:"data"`
}

type DbConfig struct {
	PGUser     string `yaml:"PG_USER"`
	PGHost     string `yaml:"PG_HOST"`
	PGPassword string `yaml:"PG_PASSWORD"`
	PGPort     string `yaml:"PG_PORT"`
	DbName     string `yaml:"DB_NAME"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}
