package cmd

import (
	"fmt"
	"github.com/ramisafariha5/daroga/contract"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(contract.GetContainerInfo().ContainerId)
		fmt.Println(contract.GetContainerInfo().ContainerName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	containerInfo := contract.GetContainerInfo()
	runCmd.Flags().StringVarP(&containerInfo.ContainerName, "cname", "n", "", "container name to inspect")
	runCmd.Flags().StringVarP(&containerInfo.ContainerId, "id", "i", "", "container id to inspect")
}
