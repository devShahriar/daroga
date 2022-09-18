package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "daroga",
	Short: "daroga",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
