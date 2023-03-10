/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/overlordtm/go-ssh-conf/pkg/ssh_conf"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	sources []string = []string{"$HOME/.ssh-conf/conf.d"}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ssh-conf",
	Short: "Create ssh_conf from partial files",
	Long:  `Create ssh_conf from partial files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		sources, err := cmd.Flags().GetStringSlice("source")
		if err != nil {
			return err
		}
		cfg, err := ssh_conf.Parse(sources)
		if err != nil {
			return err
		}
		fmt.Println(cfg)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	fmt.Println(cfgFile)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ssh-conf/conf.yaml)")
	rootCmd.PersistentFlags().StringSliceP("source", "s", sources, "Source directories/files")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
