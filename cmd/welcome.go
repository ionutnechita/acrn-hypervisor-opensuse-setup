// Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(welcomeCmd)
}

var welcomeCmd = &cobra.Command{
	Use:   "welcome",
	Short: "ACRN Hypervisor OpenSUSE Setup - Welcome",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome on Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps")
		log.Println("This project will contain instructions for installing and using the ACRN Hypervisor project.")
	},
}
