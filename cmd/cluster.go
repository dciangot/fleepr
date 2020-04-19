package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createCluster(cmd *cobra.Command, args []string) {
	fmt.Println("Creating cluster...")
}

// createCmd represents cluster create
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run:   createCluster,
}

func clusterWrapper(cmd *cobra.Command, args []string) {
	fmt.Println("Stay a while and listen...")
}

// clusterCmd represents wrapper command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "",
	Long:  ``,
	Run:   clusterWrapper,
}

func init() {
	rootCmd.AddCommand(clusterCmd)

	clusterCmd.AddCommand(createCmd)

}
