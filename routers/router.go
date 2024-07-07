package routers

import (
	"github.com/Ken-Sumi1019/samql/handlers"
	"github.com/labstack/echo/v4"
)

// 認証がないルーティングの設定を行う
//
// @param
// ctx コンテキスト
// router ルーター
func SetRouting(e *echo.Echo) {
	e.GET("/", handlers.Root)
}
