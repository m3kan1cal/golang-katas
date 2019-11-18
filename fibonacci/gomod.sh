#!/bin/bash
set -eu

touch go.mod

PROJECT_NAME=$(basename $(pwd | xargs dirname))
CURRENT_DIR=$(basename $(pwd))

CONTENT=$(cat <<-EOD
module github.com/${PROJECT_NAME}/${CURRENT_DIR}

require (
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.1
)

EOD
)

echo "$CONTENT" > go.mod
