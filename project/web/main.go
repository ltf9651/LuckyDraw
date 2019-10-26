package main

import (
	"LuckyDraw/project/bootstrap"
	"LuckyDraw/project/web/middleware/identity"
	"LuckyDraw/project/web/routes"
	"fmt"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("LuckyDraw System", "me")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf("localhost:%d", port))
}
