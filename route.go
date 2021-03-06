package catgo

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo/errors"
	"net/http"
)

// change currentRouterGroup to Router
func DefaultGroup(handlers ...interface{}) {
	handlersMap := convert2GinHandlers(handlers)
	Router.Use(handlersMap...)
	currentRouterGroup = Router
}

// create new route group
// relativePath,examples: /users /users/:id
// handlers,examples:gin.HandlerFunc,catgo.HandlerFunc
func Group(relativePath string, handlers ...interface{}) gin.IRouter {

	handlersMap := convert2GinHandlers(handlers)
	currentRouterGroup = Router.Group(relativePath, handlersMap...)

	return currentRouterGroup
}

// Handle registers a new request handle and middleware with the given path and method.
// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
// See the example code in GitHub.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func Handle(httpMethod, relativePath string, handlers ...interface{}) gin.IRouter {
	handlersMap := convert2GinHandlers(handlers)
	currentRouterGroup.Handle(httpMethod, relativePath, handlersMap...)
	return Router
}

// GET is a shortcut for catgo.Handle("GET", path, handle).
func GET(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodGet, relativePath, handlers...)
}

// POST is a shortcut for catgo.Handle("POST", path, handle).
func POST(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodPost, relativePath, handlers...)
}

// DELETE is a shortcut for catgo.Handle("DELETE", path, handle).
func DELETE(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodDelete, relativePath, handlers...)
}

// PUT is a shortcut for catgo.Handle("PUT", path, handle).
func PUT(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodPut, relativePath, handlers...)
}

// PATCH is a shortcut for catgo.Handle("PATCH", path, handle).
func PATCH(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodPatch, relativePath, handlers...)
}

// OPTIONS is a shortcut for catgo.Handle("OPTIONS", path, handle).
func OPTIONS(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodOptions, relativePath, handlers...)
}

// HEAD is a shortcut for catgo.Handle("HEAD", path, handle).
func HEAD(relativePath string, handlers ...interface{}) gin.IRouter {
	return Handle(http.MethodHead, relativePath, handlers...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func Any(relativePath string, handlers ...interface{}) gin.IRouter {
	Handle(http.MethodGet, relativePath, handlers...)
	Handle(http.MethodPost, relativePath, handlers...)
	Handle(http.MethodDelete, relativePath, handlers...)
	Handle(http.MethodPut, relativePath, handlers...)
	Handle(http.MethodPatch, relativePath, handlers...)
	Handle(http.MethodOptions, relativePath, handlers...)
	Handle(http.MethodHead, relativePath, handlers...)
	Handle(http.MethodConnect, relativePath, handlers...)
	Handle(http.MethodTrace, relativePath, handlers...)
	return Router
}

func catchHandlerErrors(context *gin.Context) {
	if err := recover(); err != nil {
		if mErr, ok := err.(errors.HttpResponseError); ok {
			switch mErr.Type {
			case "JSON":
				context.JSON(mErr.HTTPCode, mErr.Data)
			case "JSONP":
				context.JSONP(mErr.HTTPCode, mErr.Data)
			}
		}
	}
}

// convert catgo.HandlerFunc slice to gin.HandlerFunc slice
func convert2GinHandlers(handlers []interface{}) []gin.HandlerFunc {
	var handlersMap []gin.HandlerFunc

	for _, handler := range handlers {
		switch mHandler := handler.(type) {
		case HandlerFunc:
			handlersMap = append(handlersMap, func(context *gin.Context) {
				defer catchHandlerErrors(context)

				mHandler(&Context{
					Context: context,
				})
			})
		case func(*Context):
			handlersMap = append(handlersMap, func(context *gin.Context) {
				defer catchHandlerErrors(context)

				mHandler(&Context{
					Context: context,
				})
			})
		case gin.HandlerFunc:
			handlersMap = append(handlersMap, mHandler)
		case func(*gin.Context):
			handlersMap = append(handlersMap, mHandler)
		default:
			panic("unknown route handler type")
		}
	}

	return handlersMap
}
