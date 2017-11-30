package main

import (
	"fmt"

	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/logger"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()
	logger.SetEnabledConsole(true)
	//set route
	app.HttpServer.GET("/index", func(ctx dotweb.Context) error {
		_, err := ctx.WriteString("welcome to my first web!")
		return err
	})
	//begin server
	err := app.StartServer(8080)
	return err
}

func main() {
	err := StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
