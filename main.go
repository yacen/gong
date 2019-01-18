package main

import (
	"github.com/pkg/errors"
	"github.com/yacen/gong/context"
	"github.com/yacen/gong/server"
	"log"
	"time"
)

func main() {
	app := server.NewServerBuilder().
		Addr(":8000").
		ReadTimeout(10 * time.Second).
		WriteTimeout(10 * time.Second).
		MaxHeaderBytes(1 << 20).
		Build()
	app.Use(func(ctx *context.Context, chain server.Chain) {
		ctx.Res.Write([]byte("this is middleware 1, call next\n"))
		chain.Next(ctx)
	})
	app.Use(func(ctx *context.Context, chain server.Chain) {
		ctx.Res.Write([]byte("this is middleware 2, call next\n"))
		chain.Next(ctx)
	})

	app.Use(func(ctx *context.Context, chain server.Chain) {
		ctx.Res.Write([]byte("this is middleware 3\n"))
		ctx.Err = errors.New("params error")
		chain.Next(ctx)
	})

	app.Use(func(ctx *context.Context, chain server.Chain) {
		log.Println("this is error handler middleware")
		if ctx.Err != nil {
			log.Println("Error is ", ctx.Err.Error())
		}
	})
	log.Fatal(app.Listen())
}
