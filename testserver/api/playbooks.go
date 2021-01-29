package api

import "github.com/gin-gonic/gin"

var playbooks = map[string]string{
	// clustersAdded starts with a set number of clusters and adds more
	"clustersAdded":   "clusters_added",
	"clustersRemoved": "cluster_removed",
}

// RunPlaybook sets up predetermined api responses to simulate various scenarios
// for testing
func RunPlaybook(c *gin.Context) {
	name := c.Param("playbook")
	playbookFolder := playbooks[name]

}
