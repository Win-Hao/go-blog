package utils

import "new_demo/core"

var (
	AppMode      string
	HttpPort     string
	MySigningKey string

	Ip       string
	Username string
	Password string
	Port     int
	Database string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiNiuServer string
)

func init() {
	AppMode = core.Conf.Server.AppMode
	HttpPort = Str(core.Conf.Server.HttpPort)
	MySigningKey = core.Conf.Server.MySigningKey

	Ip = core.Conf.Mysql.Ip
	Username = core.Conf.Mysql.Username
	Password = core.Conf.Mysql.Password
	Port = core.Conf.Mysql.Port
	Database = core.Conf.Mysql.Database

	AccessKey = core.Conf.QiNiuYun.AccessKey
	SecretKey = core.Conf.QiNiuYun.SecretKey
	Bucket = core.Conf.QiNiuYun.Bucket
	QiNiuServer = core.Conf.QiNiuYun.QiNiuServer
}
