package database

// MongoConfig is a structure which conteins details for creating MongoDB client
type MongoConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// PostgresConfig is a structure which conteins details for creating Postgres client
type PostgresConfig struct {
	Dialect      string `json:"dialect"`
	User         string `json:"user"`
	DataBaseName string `json:"db_name"`
	SSLMode      string `json:"ssl_mode"`
	Password     string `json:"password"`
}
