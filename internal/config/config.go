package config

type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"`
}

type JWTConfig struct {
	Algorithm string `yaml:"algorithm"`
	Secret    string `yaml:"secret"`
	Issuer    string `yaml:"issuer"`
}

type Config struct {
	Database    *DatabaseConfig `yaml:"database"`
	JWT         *JWTConfig      `yaml:"jwt"`
	LogLevel    string          `yaml:"log_level"`
	ListenAddr  string          `yaml:"listen_addr"`
	SMSProvider string          `yaml:"sms_provider"`
}
