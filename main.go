package main

import "gopkg.in/kataras/iris.v4"

func main(){
	api := iris.New()
	api.Get("/", func(ctx  *iris.Context) {
		ctx.Render("home.html", struct { Name string }{ Name: "GopherHack Hyderabad" })
	})
	api.Listen(":8080")
}