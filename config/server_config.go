package config

type Server struct {
	AppMode      string `mapstructure:"AppMode"`
	HttpPort     int    `mapstructure:"HttpPort"`
	MySigningKey string `mapstructure:"mySigningKey"`
}
