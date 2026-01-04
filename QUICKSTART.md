# BlackRoad OS MinIO Quick Start Guide

## Welcome to BlackRoad OS MinIO

This guide will help you get started with BlackRoad OS MinIO, a proprietary S3-compatible object storage solution enhanced with RoadChain blockchain tracking and automated agent communication.

## Prerequisites

Before you begin, ensure you have:

1. **Valid License**: Contact licensing@blackroad-os.com to obtain a license
2. **Go 1.24+**: Required for building from source
3. **Git**: For cloning the repository
4. **API Keys**: RoadChain and Agent API keys (provided with license)

## Installation

### Step 1: Clone the Repository

```bash
git clone https://github.com/BlackRoad-OS/minio.git
cd minio
```

### Step 2: Configure Environment

Create a `.env` file with your credentials:

```bash
cat > .env << EOF
# RoadChain Configuration
ROADCHAIN_ENABLED=true
ROADCHAIN_NETWORK_URL=https://roadchain.blackroad-os.com/api/v1
ROADCHAIN_API_KEY=your-roadchain-api-key-here

# Agent Configuration
AGENT_ENABLED=true
AGENT_ID=agent-minio-$(hostname)
AGENT_ENDPOINT=https://agent.blackroad-os.com/api/v1/agents/minio
AGENT_API_KEY=your-agent-api-key-here

# MinIO Configuration
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin
EOF
```

### Step 3: Build from Source

```bash
# Load environment variables
source .env

# Build MinIO
go build -o minio

# Verify build
./minio --version
```

### Step 4: Start MinIO Server

```bash
# Create data directory
mkdir -p /tmp/minio-data

# Start server
./minio server /tmp/minio-data --console-address :9001
```

The server will start on:
- **API**: http://localhost:9000
- **Console**: http://localhost:9001

### Step 5: Verify Installation

Open your browser and navigate to http://localhost:9001

Login with:
- **Username**: minioadmin
- **Password**: minioadmin

## RoadChain Integration

### Verify RoadChain is Active

```bash
# Check RoadChain status
curl http://localhost:9000/roadchain/status

# Expected response:
# {"enabled":true,"network":"roadchain-mainnet","status":"active"}
```

### View Commit Tracking

All commits are automatically tracked via GitHub Actions. Check the Actions tab in your repository to see:

1. **RoadChain Tracking** workflow runs
2. **Agent Automation** workflow runs

### Manual Commit Tracking

You can also manually track commits using the RoadChain CLI:

```bash
# Track current commit
roadchain track --commit $(git rev-parse HEAD)

# Verify commit
roadchain verify --hash <roadchain-hash>
```

## Agent Communication

### Verify Agent Status

```bash
# Check agent status
curl http://localhost:9000/agent/status

# Expected response:
# {"agent_id":"agent-minio-yourhost","status":"online","peers":3}
```

### View Peer Agents

```bash
# List discovered peer agents
curl http://localhost:9000/agent/peers

# Expected response:
# {"peers":[
#   {"id":"agent-core","status":"online"},
#   {"id":"agent-storage","status":"online"},
#   {"id":"agent-network","status":"online"}
# ]}
```

### Send Agent Message

```bash
# Send status update to peers
curl -X POST http://localhost:9000/agent/broadcast \
  -H "Content-Type: application/json" \
  -d '{
    "message_type": "status_update",
    "payload": {
      "status": "active",
      "health": "healthy"
    }
  }'
```

## Using MinIO

### Install MinIO Client (mc)

```bash
# Download and install mc
wget https://dl.min.io/client/mc/release/linux-amd64/mc
chmod +x mc
sudo mv mc /usr/local/bin/
```

### Configure Client

```bash
# Add MinIO server
mc alias set blackroad http://localhost:9000 minioadmin minioadmin

# Verify connection
mc admin info blackroad
```

### Basic Operations

```bash
# Create a bucket
mc mb blackroad/mybucket

# Upload a file
mc cp /path/to/file.txt blackroad/mybucket/

# List files
mc ls blackroad/mybucket/

# Download a file
mc cp blackroad/mybucket/file.txt /tmp/

# Remove a file
mc rm blackroad/mybucket/file.txt
```

## Docker Deployment

### Build Docker Image

```bash
# Build the image
docker build -t blackroad-minio:latest .

# Verify image
docker images | grep blackroad-minio
```

### Run Container

```bash
# Create data volume
docker volume create minio-data

# Run container
docker run -d \
  --name blackroad-minio \
  -p 9000:9000 \
  -p 9001:9001 \
  -v minio-data:/data \
  -e ROADCHAIN_API_KEY=$ROADCHAIN_API_KEY \
  -e AGENT_API_KEY=$AGENT_API_KEY \
  blackroad-minio:latest \
  server /data --console-address :9001

# Check logs
docker logs blackroad-minio

# Check status
docker ps | grep blackroad-minio
```

## Kubernetes Deployment

### Create Namespace

```bash
kubectl create namespace blackroad-minio
```

### Create Secret

```bash
kubectl create secret generic minio-credentials \
  --from-literal=root-user=minioadmin \
  --from-literal=root-password=minioadmin \
  --from-literal=roadchain-api-key=$ROADCHAIN_API_KEY \
  --from-literal=agent-api-key=$AGENT_API_KEY \
  -n blackroad-minio
```

### Deploy

```bash
# Apply deployment
kubectl apply -f helm/minio/templates/deployment.yaml -n blackroad-minio

# Check status
kubectl get pods -n blackroad-minio

# Get service
kubectl get svc -n blackroad-minio
```

## Monitoring

### View Logs

```bash
# Application logs
tail -f /var/log/blackroad/minio.log

# RoadChain logs
tail -f /var/log/blackroad/roadchain.log

# Agent logs
tail -f /var/log/blackroad/agent.log
```

### Metrics

Access metrics endpoints:

```bash
# MinIO metrics
curl http://localhost:9000/minio/v2/metrics/cluster

# RoadChain metrics
curl http://localhost:9000/roadchain/metrics

# Agent metrics
curl http://localhost:9000/agent/metrics
```

### Health Checks

```bash
# Overall health
curl http://localhost:9000/minio/health/live

# RoadChain health
curl http://localhost:9000/roadchain/health

# Agent health
curl http://localhost:9000/agent/health
```

## Troubleshooting

### Common Issues

#### 1. License Not Valid

**Error**: "License validation failed"

**Solution**:
```bash
# Verify license key
echo $ROADCHAIN_API_KEY
echo $AGENT_API_KEY

# Contact licensing
# Email: licensing@blackroad-os.com
```

#### 2. RoadChain Not Connecting

**Error**: "Failed to connect to RoadChain network"

**Solution**:
```bash
# Check network URL
echo $ROADCHAIN_NETWORK_URL

# Test connectivity
curl -I $ROADCHAIN_NETWORK_URL

# Verify API key
curl -H "Authorization: Bearer $ROADCHAIN_API_KEY" \
  $ROADCHAIN_NETWORK_URL/status
```

#### 3. Agent Discovery Failed

**Error**: "No peer agents discovered"

**Solution**:
```bash
# Check agent endpoint
echo $AGENT_ENDPOINT

# Test discovery service
curl $AGENT_ENDPOINT/discovery

# Verify agent configuration
cat agent.config.yaml
```

#### 4. Build Failed

**Error**: "Go build failed"

**Solution**:
```bash
# Verify Go version
go version  # Should be 1.24 or higher

# Clean and rebuild
go clean -cache
go mod tidy
go build -v
```

## Getting Help

### Documentation

- **RoadChain**: See [ROADCHAIN.md](ROADCHAIN.md)
- **Agents**: See [AGENTS.md](AGENTS.md)
- **Enhancements**: See [ENHANCEMENTS.md](ENHANCEMENTS.md)
- **License**: See [LICENSE](LICENSE)

### Support

- **Email**: support@blackroad-os.com
- **Licensing**: licensing@blackroad-os.com
- **Website**: https://blackroad-os.com
- **Documentation**: https://docs.blackroad-os.com

### Reporting Issues

When reporting issues, include:

1. BlackRoad OS MinIO version
2. Operating system and version
3. Go version
4. Error messages and logs
5. Steps to reproduce
6. Configuration files (redact API keys)

## Best Practices

### Security

1. **Change Default Credentials**: Always change minioadmin/minioadmin
2. **Use TLS**: Enable TLS for production deployments
3. **Rotate Keys**: Regularly rotate API keys
4. **Network Security**: Use firewall rules and network policies
5. **Backup**: Regularly backup data and configuration

### Performance

1. **Use SSD**: Deploy on SSD storage for best performance
2. **Scale Horizontally**: Add more servers for higher throughput
3. **Monitor Resources**: Keep track of CPU, memory, and disk usage
4. **Optimize Network**: Use high-bandwidth, low-latency networks
5. **Tune Parameters**: Adjust based on workload characteristics

### Operations

1. **Monitor Continuously**: Use metrics and health checks
2. **Automate Deployments**: Use CI/CD pipelines
3. **Document Changes**: Maintain change logs
4. **Test Thoroughly**: Test in staging before production
5. **Plan Capacity**: Monitor growth and plan ahead

## Next Steps

Now that you have BlackRoad OS MinIO running:

1. ✅ **Explore the Console**: Navigate http://localhost:9001
2. ✅ **Create Buckets**: Set up your first storage buckets
3. ✅ **Upload Data**: Start storing objects
4. ✅ **Review Workflows**: Check GitHub Actions for automation
5. ✅ **Monitor Activity**: Watch RoadChain tracking and agent communication

## License

BlackRoad OS MinIO is proprietary software.

**Copyright (C) 2026 BlackRoad OS, Inc.**
All rights reserved under the BlackRoad OS Proprietary License Version 1.0.

For licensing information, contact: licensing@blackroad-os.com

---

**Thank you for choosing BlackRoad OS MinIO!**
