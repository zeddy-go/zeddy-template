package main

import (
	"github.com/zeddy-go/zeddy/app"
	_ "template/config"
	"template/module/user"

	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"github.com/zeddy-go/zeddy/http/ginx"
)

func main() {
	app.Use(
		wgorm.NewModule(),
		migrate.NewModule(),
		ginx.NewModule(),
		user.NewModule(),
	)
	err := app.StartAndWait()
	if err != nil {
		panic(err)
	}
}
