package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	inventoryCmd.AddCommand(newInventoryCmd)
}

var newInventoryCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new inventory",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts.generateHosts()
		fmt.Println("inventory new called")
	},
}
