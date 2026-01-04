// Copyright (c) 2026 BlackRoad OS, Inc.
// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of BlackRoad OS MinIO Object Storage stack
//
// This program is proprietary software: you may not redistribute it or modify
// it without explicit written permission from BlackRoad OS, Inc.
// All rights reserved under the BlackRoad OS Proprietary License Version 1.0.
//
// This program is distributed WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the BlackRoad OS Proprietary License
// along with this program. If not, contact licensing@blackroad-os.com

package main // import "github.com/BlackRoad-OS/minio"

//go:generate go install tool

import (
	"os"

	// MUST be first import.
	_ "github.com/minio/minio/internal/init"

	minio "github.com/minio/minio/cmd"
)

func main() {
	minio.Main(os.Args)
}
