#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
export COMPOSE_PROJECT_NAME="interseguro-smoke-$$"
export PUBLIC_PORT="${SMOKE_PORT:-18080}"
export MATRIX_URL="http://127.0.0.1:$PUBLIC_PORT"
trap 'docker compose down --volumes' EXIT
docker compose up --build --wait --wait-timeout 90
node scripts/smoke.mjs
docker compose stop statistics-api
node scripts/smoke.mjs --dependency-down
