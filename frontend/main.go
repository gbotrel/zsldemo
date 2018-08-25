package main

import (
	app "github.com/gbotrel/zsldemo/frontend/app"
	"github.com/gopherjs/gopherjs/js"
	vue "github.com/oskca/gopherjs-vue"
)

// JS Application objects
var (
	view        *vue.ViewModel
	application *app.App
)

//go:generate gopherjs build -m -o js/app.min.js
//go:generate echo "generated app.min.js"
//go:generate sass css/app.scss:css/app.css
//go:generate echo "generated app.css"
func main() {
	// init Vue.js app
	application = app.NewApp()
	view = vue.New(".application", application)

	// onload
	js.Global.Call("addEventListener", "load", application.OnLoad)
}
