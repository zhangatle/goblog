package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/config"
	"net/http"
)

var router *mux.Router

func main() {
	// 初始化配置信息
	config.Initialize()

	bootstrap.SetupDB()

	router = bootstrap.SetupRoute()

	http.ListenAndServe(":"+config.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
