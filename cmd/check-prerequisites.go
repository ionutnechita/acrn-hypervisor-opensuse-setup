// Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkPrerequisitesCmd)
}

var checkPrerequisitesCmd = &cobra.Command{
	Use:   "check-prerequisites",
	Short: "ACRN Hypervisor OpenSUSE Setup - Check prerequisites",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Under development")
	},
}
