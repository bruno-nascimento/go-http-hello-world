package main

import (
	"github.com/bruno-nascimento/go-http-hello-world/handlers"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/ping", handlers.Ping)
	app.Post("/new/{name:string}", handlers.New)
	app.Post("/message/{name:string}", handlers.Message)
	app.Run(iris.Addr(":8080"))
}
