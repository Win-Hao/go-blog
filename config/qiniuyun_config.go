package config

type QiNiuYun struct {
	AccessKey   string `mapstructure:"accessKey"`
	SecretKey   string `mapstructure:"secretKey"`
	Bucket      string `mapstructure:"bucket"`
	QiNiuServer string `mapstructure:"qiNiuServer"`
}
