package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"new_demo/config"
)

var Conf config.Config

func init() {
	viper.SetConfigFile("settings.yaml")
	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(fmt.Errorf("fail to read configFile:%s", err1))
	}
	err2 := viper.Unmarshal(&Conf)
	if err2 != nil {
		panic(fmt.Errorf("fail to unmarshal configFile:%s", err2))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Your have been changed" + in.Name)
	})
	err3 := viper.Unmarshal(&Conf)
	if err3 != nil {
		panic(fmt.Errorf("fail to unmarshal configFile:%s", err3))
	}
}
