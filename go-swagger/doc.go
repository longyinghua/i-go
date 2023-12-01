//go:build doc

package main

import (
	_ "go-swagger/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(files.Handler)
}