/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", versioninfo.Version)
		fmt.Println("Revision:", versioninfo.Revision)
		fmt.Println("DirtyBuild:", versioninfo.DirtyBuild)
		fmt.Println("LastCommit:", versioninfo.LastCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
