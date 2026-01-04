# BlackRoad OS Agent System

## Overview

The BlackRoad OS Agent System provides automated cross-repository communication and synchronization across all BlackRoad OS repositories. Agents enable real-time coordination, metadata sharing, and distributed automation.

## Features

- **Automatic Discovery**: Agents automatically discover peer agents in the BlackRoad OS network
- **Cross-Repository Communication**: Secure messaging between repositories
- **Real-Time Synchronization**: Metadata and status updates synced in real-time
- **Distributed Automation**: Coordinated workflows across multiple repositories
- **RoadChain Integration**: All agent communications tracked via RoadChain

## Architecture

### Components

1. **Agent Manager** (`internal/agents/agents.go`)
   - Agent registration and lifecycle management
   - Message routing and delivery
   - Peer discovery
   - Repository synchronization

2. **Configuration** (`agent.config.yaml`)
   - Agent identity and capabilities
   - Communication settings
   - Discovery configuration
   - Security parameters

3. **GitHub Actions** (`.github/workflows/agent-automation.yml`)
   - Automated agent communication
   - Peer discovery and registration
   - Status broadcasting
   - Metadata synchronization

## How It Works

1. **Agent Initialization**: Each repository has an agent that starts automatically
2. **Peer Discovery**: Agents discover other agents via the BlackRoad OS registry
3. **Registration**: Agent registers itself with central agent registry
4. **Communication**: Agents exchange messages using RoadChain RPC protocol
5. **Synchronization**: Repository metadata is synced across all peers
6. **Health Monitoring**: Agents send heartbeats and monitor peer health

## Agent Communication Protocol

### Message Format

```json
{
  "from_agent": "agent-minio-blackroad-os",
  "to_agent": "agent-core-blackroad-os",
  "message_type": "status_update",
  "payload": {
    "status": "active",
    "last_commit": "abc123",
    "branch": "main"
  },
  "timestamp": "2026-01-04T07:10:00Z",
  "signature": "sha256-hash"
}
```

### Message Types

- `status_update`: Repository status and health information
- `metadata_sync`: Repository metadata synchronization
- `commit_notification`: New commit notifications
- `deploy_request`: Deployment coordination
- `health_check`: Agent health verification

## Configuration

### Agent Configuration File

The `agent.config.yaml` file contains all agent settings:

```yaml
agent:
  id: "agent-minio-blackroad-os"
  name: "BlackRoad OS MinIO Agent"
  enabled: true
  
  communication:
    protocol: "roadchain-rpc"
    encryption: "TLS-1.3"
    endpoint: "https://agent.blackroad-os.com/api/v1/agents/minio"
    
  capabilities:
    - "storage"
    - "s3-api"
    - "roadchain-tracking"
```

### Environment Variables

- `AGENT_ENABLED`: Enable/disable agent system (default: true)
- `AGENT_ID`: Unique agent identifier
- `AGENT_ENDPOINT`: Agent communication endpoint
- `AGENT_API_KEY`: API key for agent authentication

## API Usage

### Initialize Agent Manager

```go
import "github.com/BlackRoad-OS/minio/internal/agents"

am, err := agents.InitializeAgentManager()
if err != nil {
    log.Fatal(err)
}

// Start automatic sync loop
am.StartSyncLoop()
```

### Register an Agent

```go
agent := &agents.Agent{
    ID:              "agent-custom",
    Name:            "Custom Agent",
    RepositoryURL:   "https://github.com/BlackRoad-OS/custom",
    Endpoint:        "https://agent.blackroad-os.com/api/v1/agents/custom",
    Enabled:         true,
    CommunicationProtocol: "roadchain-rpc",
}

err := am.RegisterAgent(agent)
```

### Send a Message

```go
msg := agents.Message{
    FromAgent:   "agent-minio",
    ToAgent:     "agent-core",
    MessageType: "status_update",
    Payload: map[string]interface{}{
        "status": "active",
        "commit": "abc123",
    },
    Timestamp: time.Now(),
}

err := am.SendMessage(msg)
```

### Discover Peers

```go
peers, err := am.DiscoverPeers()
if err != nil {
    log.Printf("Discovery failed: %v", err)
    return
}

fmt.Printf("Discovered %d peer agents\n", len(peers))
```

## Security

### Authentication

- **API Keys**: All agent communications require valid API keys
- **Mutual TLS**: TLS 1.3 with client certificate authentication
- **Signature Verification**: All messages cryptographically signed

### Encryption

- **Transport**: TLS 1.3 for all network communication
- **Message Payload**: AES-256-GCM encryption for sensitive data
- **Certificate Pinning**: Prevents man-in-the-middle attacks

### Authorization

- **Trusted Agents**: Only registered agents can communicate
- **Capability-Based**: Agents limited to declared capabilities
- **Rate Limiting**: Prevents abuse and DoS attacks

## Peer Repositories

The agent automatically synchronizes with these BlackRoad OS repositories:

- `BlackRoad-OS/core` - Core operating system components
- `BlackRoad-OS/storage` - Storage subsystem
- `BlackRoad-OS/network` - Network stack
- `BlackRoad-OS/security` - Security modules
- `BlackRoad-OS/*` - All BlackRoad OS repositories

## Monitoring

### Metrics

Agent system exposes the following metrics:

- `agent_messages_sent_total`: Total messages sent
- `agent_messages_received_total`: Total messages received
- `agent_peers_discovered_total`: Total peers discovered
- `agent_sync_operations_total`: Total sync operations
- `agent_errors_total`: Total errors encountered

### Health Checks

- **Endpoint**: `/health`
- **Interval**: 60 seconds
- **Timeout**: 5 seconds
- **Status**: `healthy`, `degraded`, `unhealthy`

### Logs

Agent logs are written to:
- **Location**: `/var/log/blackroad/agent.log`
- **Format**: JSON
- **Level**: Configurable (debug, info, warn, error)

## Troubleshooting

### Common Issues

1. **Agent Not Connecting**: Check `AGENT_ENDPOINT` configuration
2. **Discovery Failed**: Verify network connectivity to discovery service
3. **Authentication Error**: Ensure `AGENT_API_KEY` is valid
4. **Message Delivery Failed**: Check peer agent status

### Debug Mode

Enable debug logging:

```bash
export AGENT_LOG_LEVEL=debug
```

### Support

For technical support:
- Email: support@blackroad-os.com
- Documentation: https://docs.blackroad-os.com/agents

## Integration with RoadChain

All agent communications are tracked via RoadChain:

- **Message Tracking**: Every message is hashed and recorded
- **Verification**: Message integrity verified via blockchain
- **Audit Trail**: Complete history of agent communications
- **Provenance**: Track message origin and delivery

## Workflow Automation

Agents enable coordinated workflows across repositories:

### Example: Coordinated Deployment

```yaml
- name: Request deployment from core
  run: |
    agent send-message \
      --to agent-core \
      --type deploy_request \
      --payload '{"version": "v1.2.3", "environment": "staging"}'
```

### Example: Broadcast Status Update

```yaml
- name: Broadcast status to all peers
  run: |
    agent broadcast \
      --type status_update \
      --payload '{"status": "active", "health": "healthy"}'
```

## License

The BlackRoad OS Agent System is proprietary technology owned by BlackRoad OS, Inc.
All rights reserved under the BlackRoad OS Proprietary License Version 1.0.

Copyright (C) 2026 BlackRoad OS, Inc.
