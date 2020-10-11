package catgo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
type RouterGroup struct {
	Group gin.RouterGroup
}

func Group(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	handlersMap := convert2GinHandlers(handlers)
	currentRouterGroup = Router.Group(relativePath, handlersMap...)

	return currentRouterGroup
}

func Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRouter {
	//handlersMap := convert2GinHandlers(handlers)
	//println(currentRouterGroup)
	var handlersMap []gin.HandlerFunc

	for _, handler := range handlers {
		handlersMap = append(handlersMap, func(context *gin.Context) {
			catContext := &Context{
				Context: context,
			}
			handler(catContext)
		})
	}
	currentRouterGroup.Handle(httpMethod, relativePath, handlersMap...)
	return Router
}

func GET(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodGet, relativePath, handlers...)
}

func POST(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodPost, relativePath, handlers...)
}

func DELETE(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodDelete, relativePath, handlers...)
}

func PUT(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodPut, relativePath, handlers...)
}

func PATCH(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodPatch, relativePath, handlers...)
}

func OPTIONS(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodOptions, relativePath, handlers...)
}

func HEAD(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	return Handle(http.MethodHead, relativePath, handlers...)
}

func Any(relativePath string, handlers ...HandlerFunc) gin.IRouter {
	Handle(http.MethodGet, relativePath, handlers...)
	Handle(http.MethodPost, relativePath, handlers...)
	Handle(http.MethodDelete, relativePath, handlers...)
	Handle(http.MethodPut, relativePath, handlers...)
	Handle(http.MethodPatch, relativePath, handlers...)
	Handle(http.MethodOptions, relativePath, handlers...)
	Handle(http.MethodHead, relativePath, handlers...)

	return Router
}

func convert2GinHandlers(handlers []HandlerFunc) []gin.HandlerFunc {
	var handlersMap []gin.HandlerFunc

	for _, handler := range handlers {
		handlersMap = append(handlersMap, func(context *gin.Context) {
			catContext := &Context{
				Context: context,
			}
			handler(catContext)
		})
	}

	return handlersMap
}
