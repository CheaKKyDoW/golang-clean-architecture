package config

import (
	"time"
)

type AppConfig struct {
	Name            string        `env:"NAME"`
	ENV             string        `env:"ENV" envDefault:"local"`
	Host            string        `env:"HOST" envDefault:"localhost"`
	Port            uint16        `env:"PORT" envDefault:"3000"`
	Debug           bool          `env:"DEBUG" envDefault:"true"`
	ClientTimeout   time.Duration `env:"CLIENT_TIMEOUT" envDefault:"90s"`
	RateLimitConn   int           `env:"RATE_LIMIT_CONN" envDefault:"1000"`
	RateLimitWindow time.Duration `env:"RATE_LIMIT_WINDOW" envDefault:"10s"`
}

type JWTConfig struct {
	Secret   string        `env:"SECRET,notEmpty"`
	Expire   time.Duration `env:"EXPIRE" envDefault:"1h"`
	Issuer   string        `env:"ISSUER,notEmpty"`
	Audience []string      `env:"AUDIENCE,notEmpty"`
}

type DatabaseConfig struct {
	Driver     string        `env:"DRIVER"`
	Host       string        `env:"HOST" envDefault:"localhost"`
	Port       uint16        `env:"PORT" envDefault:"5432"`
	Username   string        `env:"USER,notEmpty"`
	Password   string        `env:"PASSWORD,notEmpty"`
	Name       string        `env:"NAME,notEmpty"`
	SearchPath string        `env:"SEARCH_PATH,notEmpty" envDefault:"public"`
	TZ         string        `env:"TZ" envDefault:"UTC"`
	SSLMode    string        `env:"SSL_MODE" envDefault:"disable"`
	MaxOpen    uint16        `env:"MAX_OPEN" envDefault:"10"`
	MaxIdle    uint16        `env:"MAX_IDLE" envDefault:"2"`
	MaxLife    time.Duration `env:"MAX_LIFE" envDefault:"1h"`
}

type RedisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint16 `env:"PORT" envDefault:"6379"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD,notEmpty"`
	Name     string `env:"NAME" envDefault:"0"`
}

type Cfg struct {
	App      AppConfig      `env:",file" envPrefix:"APP_"`
	Database DatabaseConfig `env:",file" envPrefix:"DATABASE_"`
	JWT      JWTConfig      `env:",file" envPrefix:"JWT_"`
}
