package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-cluster-management/discovery/pkg/api/domain/cluster_domain"
)

// GetCluster ...
func GetCluster(c *gin.Context) {
	clusterID := c.Param("clusterID")
	params := c.Request.URL.Query()

	var file []byte
	var err error
	// Return filtered results if search param set
	if _, ok := params["search"]; ok {
		file, err = ioutil.ReadFile("data/filtered_clusters_list.json")
	} else {
		file, err = ioutil.ReadFile("data/clusters_list.json")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
		})
		return
	}

	var clusterList cluster_domain.ClusterResponse
	err = json.Unmarshal(file, &clusterList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()),
		})
		return
	}

	if clusterID != "" {
		c.JSON(http.StatusOK, clusterList)
	}

	for _, cluster := range clusterList.Items {
		if cluster.ID == clusterID {
			c.JSON(http.StatusOK, cluster)
			return
		}
	}

	c.Status(http.StatusBadRequest)
}
