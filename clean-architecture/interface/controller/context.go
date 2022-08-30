package controller

type Context interface {
	Param(key string) string
	JSON(code int, obj interface{})
	Abort()
}
