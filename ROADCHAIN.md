# RoadChain Integration for BlackRoad OS MinIO

## Overview

RoadChain is BlackRoad OS's proprietary blockchain-based commit tracking and verification system. It provides SHA-256 cryptographic hashing of all commits with immutable blockchain storage for complete code provenance and integrity verification.

## Features

- **SHA-256 Commit Tracking**: Every commit is hashed using SHA-256 and recorded on the RoadChain blockchain
- **Blockchain Verification**: All commits are verified against the blockchain for integrity
- **Immutable Audit Trail**: Complete history of all code changes stored permanently on the blockchain
- **Cross-Repository Provenance**: Track code lineage across all BlackRoad OS repositories
- **Automated Tracking**: GitHub Actions workflows automatically track commits in real-time

## Architecture

### Components

1. **RoadChain Core** (`internal/roadchain/roadchain.go`)
   - SHA-256 hash computation
   - Blockchain integration
   - Commit verification
   - History retrieval

2. **Configuration** (`roadchain.config.yaml`)
   - RoadChain network settings
   - Blockchain configuration
   - Verification parameters
   - Security settings

3. **GitHub Actions** (`.github/workflows/roadchain-tracking.yml`)
   - Automatic commit tracking
   - Hash computation on push
   - Blockchain record creation
   - Peer repository synchronization

## How It Works

1. **Commit Creation**: When a developer creates a commit, Git records the change locally
2. **Push to GitHub**: The commit is pushed to the GitHub repository
3. **GitHub Action Trigger**: The push triggers the RoadChain tracking workflow
4. **Hash Computation**: The workflow computes a SHA-256 hash of the commit metadata
5. **Blockchain Recording**: The hash is submitted to the RoadChain blockchain network
6. **Verification**: The commit is verified and confirmed on the blockchain
7. **Peer Sync**: The commit information is synchronized with peer repositories

## Configuration

### RoadChain Configuration File

The `roadchain.config.yaml` file contains all RoadChain settings:

```yaml
roadchain:
  enabled: true
  version: "1.0.0"
  
  commit_tracking:
    enabled: true
    hash_algorithm: "sha256"
    verify_signatures: true
    blockchain_integration: true
    
  blockchain:
    network: "roadchain-mainnet"
    node_url: "https://roadchain.blackroad-os.com/api/v1"
    consensus_algorithm: "proof-of-authority"
```

### Environment Variables

- `ROADCHAIN_ENABLED`: Enable/disable RoadChain tracking (default: true)
- `ROADCHAIN_NETWORK_URL`: RoadChain blockchain node URL
- `ROADCHAIN_API_KEY`: API key for RoadChain services (required for production)

## API Usage

### Initialize RoadChain

```go
import "github.com/BlackRoad-OS/minio/internal/roadchain"

rc, err := roadchain.InitializeRoadChain()
if err != nil {
    log.Fatal(err)
}
```

### Track a Commit

```go
commit := roadchain.CommitRecord{
    Author:    "developer@blackroad-os.com",
    Message:   "Add new feature",
    Timestamp: time.Now(),
}

err := rc.TrackCommit(commit)
if err != nil {
    log.Printf("Failed to track commit: %v", err)
}
```

### Verify a Commit

```go
hash := "a1b2c3d4e5f6..."
verified, err := rc.VerifyCommit(hash)
if err != nil {
    log.Printf("Verification failed: %v", err)
}

if verified {
    fmt.Println("Commit verified successfully")
}
```

## Security

### Cryptographic Hash Function

RoadChain uses SHA-256 for all commit hashing:
- **Algorithm**: SHA-256 (Secure Hash Algorithm 256-bit)
- **Output**: 256-bit (32-byte) hash value
- **Collision Resistance**: Cryptographically secure
- **Performance**: Optimized for high-throughput commit tracking

### Blockchain Security

- **Consensus**: Proof-of-Authority (PoA)
- **Network**: Private BlackRoad OS blockchain network
- **Encryption**: TLS 1.3 for all communications
- **Authentication**: API key-based authentication required

## Monitoring

### Metrics

RoadChain exposes the following metrics:

- `roadchain_commits_tracked_total`: Total number of commits tracked
- `roadchain_commits_verified_total`: Total number of verified commits
- `roadchain_verification_failures_total`: Total verification failures
- `roadchain_sync_duration_seconds`: Duration of peer synchronization

### Logs

RoadChain logs are written to:
- **Location**: `/var/log/blackroad/roadchain.log`
- **Format**: JSON
- **Level**: Configurable (debug, info, warn, error)

## Troubleshooting

### Common Issues

1. **Tracking Disabled**: Ensure `ROADCHAIN_ENABLED=true` in configuration
2. **Network Errors**: Check `ROADCHAIN_NETWORK_URL` is accessible
3. **Authentication Failed**: Verify `ROADCHAIN_API_KEY` is valid
4. **Verification Failed**: Check blockchain network status

### Support

For technical support:
- Email: support@blackroad-os.com
- Documentation: https://docs.blackroad-os.com/roadchain

## License

RoadChain is proprietary technology owned by BlackRoad OS, Inc.
All rights reserved under the BlackRoad OS Proprietary License Version 1.0.

Copyright (C) 2026 BlackRoad OS, Inc.
