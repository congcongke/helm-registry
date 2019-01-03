/*
Copyright 2017 congcongke authors. All rights reserved.
*/

package cmd

import (
	"github.com/congcongke/helm-registry/pkg/log"
	"github.com/spf13/cobra"
)

// rootCmd is a root command and shows help information
var rootCmd = &cobra.Command{
	Use:   "registry",
	Short: "registry is a repository of helm charts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Run executes rootCmd
func Run() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
