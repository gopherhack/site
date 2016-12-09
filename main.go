package main

import (
	"gopkg.in/kataras/iris.v4"
	"os"
	"github.com/kataras/go-template/html"
)

func main(){

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	api := iris.New()
	api.UseTemplate(html.New()).Directory("./public/templates", ".html")
	api.Get("/", func(ctx  *iris.Context) {
		ctx.Render("home.html", struct { Name string }{ Name: "GopherHack Hyderabad" })
	})
	api.Listen(":"+port)
}