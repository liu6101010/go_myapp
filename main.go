package main

import (
	"fmt"
	"myapp/libs"

	"github.com/kataras/iris/v12"
)

var l_p_a int
var l_p_b string

func show_page(x int, y int) (r1 int, r2 int) {
	fmt.Println("######## x=", x, "y=", y, "\n")
	//var z1 = 11
	//var z2 = 12
	r1 = 22
	r2 = 33
	l_p_a, l_p_b = libs.GetParams()
	libs.Show()
	fmt.Println("######## l_p_a=", l_p_a, "l_p_b=", l_p_b)
	return
}

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("######## Hello <strong>%d %s</strong>!World ", l_p_a, l_p_b)
	})
	fmt.Println("######## iris running \n")
	z1, z2 := show_page(9, 10)
	fmt.Println("######## z1=[", z1, "] z2=[", z2, "]")

	app.Listen(":8080")
}
