package config

type Mysql struct {
	Ip       string `mapstructure:"ip"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
}
