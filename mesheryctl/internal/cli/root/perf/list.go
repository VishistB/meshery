package perf

import (
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Performance profiles",
	Long:    `List all the available performance profiles`,
	Example: "mesheryctl perf list",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get viper instance used for context
		// mctlCfg, err := config.GetMesheryCtl(viper.GetViper())
		// if err != nil {
		// 	return errors.Wrap(err, "error processing config")
		// }

		var data [][]string
		utils.PrintToTable([]string{"NAME", "RESULTS", "LAST-RUN"}, data)
		return nil
	},
}
