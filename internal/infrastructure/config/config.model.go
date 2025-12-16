package config

type ConfigModel struct {
	AppName    string `bind:"APP_NAME" validate:"required"`
	AppEnv     string `bind:"APP_ENV" validate:"required"`
	ServerPort int    `bind:"SERVER_PORT" validate:"required"`

	JWTAccessAPIKey      string `bind:"JWT_ACCESS_API_KEY" validate:"required"`
	JWTAccessAPIExpIn    int    `bind:"JWT_ACCESS_API_EXP_IN" validate:"required"`
	JWTRefreshTokenKey   string `bind:"JWT_REFRESH_TOKEN_KEY" validate:"required"`
	JWTRefreshTokenExpIn int    `bind:"JWT_REFRESH_TOKEN_EXP_IN" validate:"required"`

	PostgresHost     string `bind:"POSTGRES_HOST" validate:"required"`
	PostgresUser     string `bind:"POSTGRES_USER" validate:"required"`
	PostgresPassword string `bind:"POSTGRES_PASSWORD" validate:"required"`
	PostgresDB       string `bind:"POSTGRES_DB" validate:"required"`
	PostgresPort     int    `bind:"POSTGRES_PORT" validate:"required"`
	PostgresSSL      bool   `bind:"POSTGRES_SSL" validate:"boolean"`

	DomainWhitelist []string `bind:"DOMAIN_WHITELIST" validate:"required"`
}
