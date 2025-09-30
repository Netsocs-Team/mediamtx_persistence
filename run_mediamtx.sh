#!/usr/bin/env bash
set -euo pipefail

# Default paths
MEDIAMTX_BINARY=${MEDIAMTX_BINARY:-./mediamtx}
CONFIG_PATH=${CONFIG_PATH:-./config/mediamtx.yaml}
API_URL=${API_URL:-http://localhost:9997}

# Ensure directories exist
mkdir -p recordings
mkdir -p config

echo "Starting MediaMTX with persistence..."
echo "Binary: $MEDIAMTX_BINARY"
echo "Config: $CONFIG_PATH"
echo "API URL: $API_URL"

# Run the persistence binary which will start MediaMTX
exec ./persistence \
    --mediamtx-binary="$MEDIAMTX_BINARY" \
    --config="$CONFIG_PATH" \
    --mediamtx="$API_URL"
