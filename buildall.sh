#!/usr/bin/env bash
set -euo pipefail

NAME="palcli"
OUTDIR="dist"

mkdir -p "$OUTDIR"

platforms=(
  "linux/amd64"
  "linux/arm64"
  "linux/arm/7,hardfloat"
  "windows/amd64"
)

for platform in "${platforms[@]}"; do
  IFS="/" read -r GOOS GOARCH GOARM <<< "$platform"

  output="$OUTDIR/${NAME}-${GOOS}-${GOARCH}"
  [ "$GOOS" = "windows" ] && output+=".exe"

  echo "Building $output"

  if [ -n "${GOARM:-}" ]; then
    GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM CGO_ENABLED=0 \
      go build -ldflags="-s -w" -o "$output"
  else
    GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 \
      go build -ldflags="-s -w" -o "$output"
  fi
done

echo "Done. Outputs in $OUTDIR/"

