package instance

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/valyentdev/ravel/pkg/core/api"
)

func GetClient(cmd *cobra.Command) *api.AgentClient {
	return api.NewAgentClient(http.DefaultClient, "http://127.0.0.1:8080")
}
