package catgo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRouter {
	var handlersMap []gin.HandlerFunc

	for _, handler := range handlers {
		handlersMap = append(handlersMap, func(context *gin.Context) {
			catContext := &Context{
				Context: context,
			}
			handler(catContext)
		})
	}
	Router.Handle(httpMethod, relativePath, handlersMap...)
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