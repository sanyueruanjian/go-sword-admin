package middleware

import (
	sentinelPlugin "github.com/alibaba/sentinel-golang/adapter/gin"
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Sentinel 限流
func Sentinel(triggerCount float64) gin.HandlerFunc {
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: triggerCount,
			Strategy:     system.BBR,
		},
	}); err != nil {
		zap.L().Fatal("Unexpected error", zap.Error(err))
	}
	return sentinelPlugin.SentinelMiddleware()
}
