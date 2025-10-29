package config

type Config struct {
	Env string `konaf: "env" validate:"required"`
}

type ServerConfig struct {
	Port               string   `konaf:"port" validate:"required"`
	ReadTimeout        int      `konaf:"read_timeout" validate:"required"`
	WriteTimeout       int      `konaf: "write_timeout" validate:"required"`
	IdleTimeout        int      `konaf: "idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `konaf: "cors_allowed_origins" validate:"required"`
}
