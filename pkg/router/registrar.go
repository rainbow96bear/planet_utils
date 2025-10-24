package router

import "github.com/gin-gonic/gin"

// RouteRegistrar는 각 핸들러가 Routes를 등록할 때 구현
type RouteRegistrar interface {
	RegisterRoutes(r *gin.Engine)
}
