# BlackRoad OS - Proprietary Edition

**This project is proprietary software owned by BlackRoad OS, Inc.**

- The codebase is proprietary and all rights are reserved
- Usage requires a valid license from BlackRoad OS, Inc.
- Contact licensing@blackroad-os.com for licensing information
- Powered by RoadChain technology for secure SHA-256 commit tracking
- Enhanced with automated agents for cross-repository communication

For enterprise licensing and support, please contact [BlackRoad OS, Inc.](https://blackroad-os.com/).

---

# BlackRoad OS MinIO Quickstart Guide

[![license](https://img.shields.io/badge/license-Proprietary-red)](https://github.com/BlackRoad-OS/minio/blob/master/LICENSE)

[![BlackRoad OS](https://raw.githubusercontent.com/minio/minio/master/.github/logo.svg?sanitize=true)](https://blackroad-os.com)

BlackRoad OS MinIO is a high-performance, S3-compatible object storage solution released under the BlackRoad OS Proprietary License.
Designed for speed and scalability, it powers AI/ML, analytics, and data-intensive workloads with industry-leading performance.

- S3 API Compatible – Seamless integration with existing S3 tools
- Built for AI & Analytics – Optimized for large-scale data pipelines
- High Performance – Ideal for demanding storage workloads
- RoadChain Enabled – SHA-256 commit tracking with blockchain verification
- Agent Automated – Cross-repository communication and automation

This README provides instructions for building BlackRoad OS MinIO from source and deploying onto baremetal hardware.

## BlackRoad OS MinIO is Proprietary Software

BlackRoad OS MinIO is proprietary software owned and licensed by BlackRoad OS, Inc. 
All usage requires a valid license agreement with BlackRoad OS, Inc.

The BlackRoad OS Proprietary License provides no rights to copy, modify, distribute, 
or use this software without explicit written permission from BlackRoad OS, Inc.

All support is provided through commercial licensing agreements. For more information, 
contact licensing@blackroad-os.com or visit https://blackroad-os.com/.

## RoadChain Integration

BlackRoad OS MinIO integrates RoadChain technology for secure commit tracking:
- SHA-256 cryptographic hashing of all commits
- Blockchain-based verification for code integrity
- Immutable audit trail of all code changes
- Cross-repository provenance tracking

## Proprietary Distribution

**Important:** BlackRoad OS MinIO is distributed as proprietary software. All usage requires a valid license from BlackRoad OS, Inc.

### Installing BlackRoad OS MinIO

To use BlackRoad OS MinIO, you must have a valid license. Contact licensing@blackroad-os.com for licensing options.

## Build from Source (Licensed Users Only)

Use the following commands to compile and run BlackRoad OS MinIO from source.
**Note:** Building requires a valid license. Ensure you have received authorization from BlackRoad OS, Inc.

If you do not have a working Golang environment, please follow [How to install Golang](https://golang.org/doc/install). Minimum version required is [go1.24](https://golang.org/dl/#stable)

```sh
go install github.com/BlackRoad-OS/minio@latest
```

You can alternatively run `go build` and use the `GOOS` and `GOARCH` environment variables to control the OS and architecture target.
For example:

```
env GOOS=linux GOARCH=arm64 go build
```

Start BlackRoad OS MinIO by running `minio server PATH` where `PATH` is any empty folder on your local filesystem.

The BlackRoad OS MinIO deployment starts using default root credentials `minioadmin:minioadmin`.
You can test the deployment using the MinIO Console, an embedded web-based object browser built into MinIO Server.
Point a web browser running on the host machine to <http://127.0.0.1:9000> and log in with the root credentials.
You can use the Browser to create buckets, upload objects, and browse the contents of the MinIO server.

You can also connect using any S3-compatible tool, such as the MinIO Client `mc` commandline tool:

```sh
mc alias set local http://localhost:9000 minioadmin minioadmin
mc admin info local
```

See [Test using MinIO Client `mc`](#test-using-minio-client-mc) for more information on using the `mc` commandline tool.

> [!NOTE]
> Production environments using BlackRoad OS MinIO require a valid commercial license.
> Contact licensing@blackroad-os.com for enterprise licensing options.

## Build Docker Image (Licensed Users Only)

You can use the `docker build .` command to build a Docker image on your local host machine.
You must first [build BlackRoad OS MinIO](#build-from-source-licensed-users-only) and ensure the `minio` binary exists in the project root.

The following command builds the Docker image using the default `Dockerfile` in the root project directory with the repository and image tag `blackroad-minio:latest`

```sh
docker build -t blackroad-minio:latest .
```

Use `docker image ls` to confirm the image exists in your local repository.
You can run the server using standard Docker invocation:

```sh
docker run -p 9000:9000 -p 9001:9001 blackroad-minio:latest server /tmp/minio --console-address :9001
```

Complete documentation for building Docker containers, managing custom images, or loading images into orchestration platforms is out of scope for this documentation.
You can modify the `Dockerfile` and `dockerscripts/docker-entrypoint.sh` as-needed to reflect your specific image requirements.

## Install using Helm Charts (Licensed Users Only)

BlackRoad OS MinIO can be deployed onto Kubernetes infrastructure using Helm charts.
**Note:** Kubernetes deployment requires a valid enterprise license.

Contact BlackRoad OS, Inc. for Kubernetes deployment options and licensing.

## Test BlackRoad OS MinIO Connectivity

### Test using MinIO Console

BlackRoad OS MinIO Server comes with an embedded web based object browser.
Point your web browser to <http://127.0.0.1:9000> to ensure your server has started successfully.

> [!NOTE]
> BlackRoad OS MinIO runs console on random port by default, if you wish to choose a specific port use `--console-address` to pick a specific interface and port.

### Test using MinIO Client `mc`

`mc` provides a modern alternative to UNIX commands like ls, cat, cp, mirror, diff etc. It supports filesystems and Amazon S3 compatible cloud storage services.

The following commands set a local alias, validate the server information, create a bucket, copy data to that bucket, and list the contents of the bucket.

```sh
mc alias set local http://localhost:9000 minioadmin minioadmin
mc admin info
mc mb data
mc cp ~/Downloads/mydata data/
mc ls data/
```

Follow the MinIO Client Quickstart Guide for further instructions.

## Explore Further

- Contact licensing@blackroad-os.com for licensing information
- Visit https://blackroad-os.com/ for enterprise support

## Contribute to BlackRoad OS MinIO Project

This is proprietary software. Contributions require a contributor license agreement (CLA).
Please contact legal@blackroad-os.com for contribution guidelines.

## License

- BlackRoad OS MinIO source is licensed under the [BlackRoad OS Proprietary License](https://github.com/BlackRoad-OS/minio/blob/master/LICENSE).
- All rights reserved. Copyright (C) 2026 BlackRoad OS, Inc.
- RoadChain integration for secure SHA-256 commit tracking included.
