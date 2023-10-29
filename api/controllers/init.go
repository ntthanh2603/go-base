package controllers

import (
	"go-base/driver"

	"github.com/gin-gonic/gin"
)

type ControllerBase struct {
	basePath string
	routes   []RouteBase
}

func Controller(params ControllerBase) *gin.Engine {
	var drive = driver.Instance
	c := drive.Group(params.basePath)
	for _, route := range params.routes {
		switch route.method {
		case "GET":
			c.GET(route.basePath, route.handler)
		case "POST":
			c.POST(route.basePath, route.handler)
		case "PUT":
			c.PUT(route.basePath, route.handler)
		case "DELETE":
			c.DELETE(route.basePath, route.handler)
		case "PATCH":
			c.PATCH(route.basePath, route.handler)
		}
	}
	return drive
}

type RouteBase struct {
	basePath string
	handler  gin.HandlerFunc
	method   string
}

type MethodHandler struct {
	basePath string
	handler  gin.HandlerFunc
	configs  MethodHandlerConfigs
}

type MethodHandlerConfigs struct {
	openApi OpenApiConfigs
}

type OpenApiConfigs struct {
	title       string
	description string
	version     string
	tags        []string
	success     string
}

func Get(basePath string, handler gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   "GET",
	}
}

func Post(basePath string, handler gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   "POST",
	}
}

func Put(basePath string, handler gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   "PUT",
	}
}

func Delete(basePath string, handler gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   "DELETE",
	}
}

func Patch(basePath string, handler gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   "PATCH",
	}
}
