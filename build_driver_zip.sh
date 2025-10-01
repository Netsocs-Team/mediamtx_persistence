#!/usr/bin/env bash

set -euo pipefail

# Configuration
PACKAGE_NAME=${PACKAGE_NAME:-mediamtx_with_persistence.zip}
MEDIA_MTX_URL=${MEDIA_MTX_URL:-"https://netsocs-public-drivers-repository.s3.us-east-1.amazonaws.com/video_engine_mediamtx/mediamtx"}

echo "Package name: ${PACKAGE_NAME}"
echo "MediaMTX URL: ${MEDIA_MTX_URL}"

# Check requirements
for cmd in curl zip go; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
        echo "Error: required command '$cmd' not found in PATH" >&2
        exit 1
    fi
done

# Build persistence binary (static-ish)
echo "Building persistence binary..."
export CGO_ENABLED=0
go build -trimpath -ldflags "-s -w" 

# Fetch MediaMTX binary
echo "Downloading MediaMTX binary..."
curl -fL "$MEDIA_MTX_URL" -o ./mediamtx
chmod +x ./mediamtx
./mediamtx --version || true

# Prepare directories to include in package
echo "Preparing directories..."
mkdir -p recordings
touch recordings/.gitkeep

# Create zip
echo "Creating zip package '${PACKAGE_NAME}'..."
rm -f "$PACKAGE_NAME"
zip -r "$PACKAGE_NAME" driver.mediamtx_persistence mediamtx recordings driver.netsocs.json mediamtx.yml
echo "Done. Wrote $(pwd)/${PACKAGE_NAME}"


