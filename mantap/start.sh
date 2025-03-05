#!/bin/sh

set -e

echo "run db migration"
migrate -path migration -database "${DATABASE_URL}" -verbose up 2

echo "start the app"
exec "$@"