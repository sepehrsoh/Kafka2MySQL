package providers

import (
	"go.uber.org/zap"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func NewGinServer() *gin.Engine {
	engine := gin.New()
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use()
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	return engine
}
