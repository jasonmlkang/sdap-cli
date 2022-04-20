/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var ingestionCmd = &cobra.Command{
	Use:   "ingestion",
	Short: "SDAP ingestion",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("ingestion called")
	},
}

func init() {
	rootCmd.AddCommand(ingestionCmd)
}
