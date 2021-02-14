# Copyright (c) 2021, Project ACRN Embedded Hypervisor using OpenSUSE OS - Setup Steps, Ionut Nechita.
# All rights reserved.
# SPDX-License-Identifier: BSD-3-Clause

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required(VERSION 3.5)

project(acrn-hypervisor-opensuse-setup Go)

include(cmake/golang.cmake)
include(cmake/flags.cmake)

set(SOURCE_FILE main.go)
add_go_executable(acrn-hypervisor-opensuse-setup ${SOURCE_FILE})
