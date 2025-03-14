package engines

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinEnginWrapper struct {
	ginEngine *gin.Engine
}

func (g *GinEnginWrapper) Handler() http.Handler {
	return g.ginEngine
}

func Gin(g *gin.Engine) *GinEnginWrapper {
	return &GinEnginWrapper{ginEngine: g}
}
