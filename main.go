package main

import (
	_ "new_demo/docs"
	"new_demo/routers"
)

// @title 规范地写一个go_demo
// @version 1.0
// @description 从头开始写一个demo
// @contact.name 只因哥
// @contact.email www.758494478@qq.com
// @host localhost:8093
// @BasePath /
func main() {
	routers.InitRouter()
}
