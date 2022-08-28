// Package probe
// @Description:
package probe

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {
	router.GET("healthy", Healthy)
}
