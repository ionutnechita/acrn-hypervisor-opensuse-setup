// Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"../pkg/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "ACRN Hypervisor OpenSUSE Setup - View application information",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(app.GetVERSIONAPP())
	},
}
