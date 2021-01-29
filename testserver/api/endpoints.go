package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// SetupEndpoints ...
func SetupEndpoints(r *gin.Engine, logger zerolog.Logger) {
	r.GET("/playbooks/*playbook", RunPlaybook)
	r.GET("/api/clusters_mgmt/v1/clusters/*clusterID", GetCluster)
	r.POST("/auth/realms/redhat-external/protocol/openid-connect/token", GetToken)
	// r.GET("/api/accounts_mgmt/v1/subscriptions", GetSubscriptions)
}
