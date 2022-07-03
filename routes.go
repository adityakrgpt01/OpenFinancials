package main

import(
	"github.com/labstack/echo"
)

func AddRoute(e *echo.Echo){
	e.GET("bitly/{shortedUrl}",func(ctx echo.Context) error {
		return GetShortedUrl(req)
	})

	e.POST("bitly/{UserId}/ActualUrl",func(ctx echo.Context) error {
		return AddUrl(req);
	})
}
