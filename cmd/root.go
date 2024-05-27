/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/overlordtm/go-ssh-conf/pkg/ssh_conf"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	outFilePath string
	sources     []string = []string{"$HOME/.ssh-conf/conf.d"}
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
			fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		}

		var outFile io.Writer

		if outFilePath != "" {
			outFile, err = os.OpenFile(outFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to open output file: %v\n", err)
				return err
			}
		} else {
			outFile = os.Stdout
		}

		_, err = io.Copy(outFile, bytes.NewBufferString(cfg.String()))
		return err
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ssh-conf/conf.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outFilePath, "out", "o", "", "output file, defaults to stdout")
	rootCmd.PersistentFlags().StringSliceP("source", "s", sources, "Source directories/files")
}
