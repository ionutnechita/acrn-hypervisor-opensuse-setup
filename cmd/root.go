// Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acrn-hypervisor-opensuse-setup",
	Short: "ACRN Hypervisor OpenSUSE Setup",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help(); os.Exit(0) },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
