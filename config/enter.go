package config

type Config struct {
	Mysql    Mysql    `mapstructure:"mysql"`
	Server   Server   `mapstructure:"server"`
	QiNiuYun QiNiuYun `mapstructure:"qiNiuYun"`
}
