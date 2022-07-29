package main

import (
	auth "kumu-exam/src/authenticator"
	git "kumu-exam/src/controllers/git"
	hlpr "kumu-exam/src/helpers"
	mdlw "kumu-exam/src/middlewares"
	"os"

	"github.com/kataras/iris/v12"
)

func init() {
	hlpr.InitCache()
}

func initApp() *iris.Application {
	app := iris.New()

	// Root
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": ctx.GetStatusCode(), "message": "Hi! This is KUMU Coding Challenge"})
	})

	// Public Endpoints
	token := app.Party("/auth")
	{
		token.Handle("POST", "/token", auth.AuthToken)
	}

	// Authenticated Endpoints
	api := app.Party("/api/v1", mdlw.ValidateToken)
	{
		api.Handle("GET", "/git-users", git.GetUsers)
	}

	return app
}

func main() {
	app := initApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "1007"
	}
	app.Listen(":" + port)
}
