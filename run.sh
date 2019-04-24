#!/usr/bin/env bash

set -e

repo_dir="$( cd "$( dirname "$0" )" && pwd )"

(
    cd "${repo_dir}"
    go generate ./...
    go run . solve "$@"
)
