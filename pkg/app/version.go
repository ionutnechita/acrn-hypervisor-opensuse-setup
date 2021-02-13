// Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package app

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

// VERSION as const
const VERSION = "0"

// PATCHLEVEL as const
const PATCHLEVEL = "1"

// SUBLEVEL as const
const SUBLEVEL = "7"

// EXTRALEVEL as const
const EXTRALEVEL = "-DEV"

// NAME as const
const NAME = "Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps - General Purpose Flavor"

// GetVERSION as function, return string, error
func GetVERSION() (string, error) {
	if VERSION != "" {
		msg := fmt.Sprintf(VERSION)
		return msg, nil
	}
	return "", errors.New("Version not inserted in program")
}

// GetPATCHLEVEL as function, return string, error
func GetPATCHLEVEL() (string, error) {
	if PATCHLEVEL != "" {
		msg := fmt.Sprintf("." + PATCHLEVEL)
		return msg, nil
	}
	return "", errors.New("Patch level not inserted in program. Maybe is optional")
}

// GetSUBLEVEL as function, return string, error
func GetSUBLEVEL() (string, error) {
	if SUBLEVEL != "" {
		msg := fmt.Sprintf("." + SUBLEVEL)
		return msg, nil
	}
	return "", errors.New("Sub level not inserted in program. Maybe is optional")
}

// GetEXTRALEVEL as function, return string, error
func GetEXTRALEVEL() (string, error) {
	if EXTRALEVEL != "" {
		msg := fmt.Sprintf(EXTRALEVEL)
		return msg, nil
	}
	return "", errors.New("Extra level not inserted in program. Maybe is optional")
}

// GetNAME as function, return string, error
func GetNAME() (string, error) {
	if NAME != "" {
		msg := fmt.Sprintf(NAME)
		return msg, nil
	}
	return "", errors.New("Application name not inserted in program")
}

// GetVERSIONAPP as function, return string
func GetVERSIONAPP() string {
	version, err := GetVERSION()
	if err != nil {
		log.Fatal(err)
	}
	patchlevel, err := GetPATCHLEVEL()
	if err != nil {
		log.Print(err)
	}
	sublevel, err := GetSUBLEVEL()
	if err != nil {
		log.Print(err)
	}
	extralevel, err := GetEXTRALEVEL()
	if err != nil {
		log.Print(err)
	}
	name, err := GetNAME()
	if err != nil {
		log.Fatal(err)
	}
	nameApplication := name
	versionApplication := version + patchlevel + sublevel + extralevel
	timeNowApplication := time.Now().Format("2006-01-02 15:04:05")
	osPlatformApplication := runtime.GOOS
	osArchApplication := runtime.GOARCH
	msgNameApplication := fmt.Sprintf("\nNAME: " + nameApplication)
	msgVersionApplication := fmt.Sprintf("\nVERSION: " + versionApplication)
	msgOsPlatformApplication := fmt.Sprintf("\nOS PLATFORM: " + osPlatformApplication)
	msgOsArchApplication := fmt.Sprintf("\nOS ARCH: " + osArchApplication)
	msgTimeNowApplication := fmt.Sprintf("\nTIME NOW: " + timeNowApplication)
	return msgNameApplication +
		msgVersionApplication +
		msgOsPlatformApplication +
		msgOsArchApplication +
		msgTimeNowApplication
}
