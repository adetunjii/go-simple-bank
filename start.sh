#!/bin/sh

# exit immediately if there's an error
set -e
echo "running db migration...."
/app/migrate - path /app/migration -database "$DB_SOURCE" -verbose

echo "starting the app 🚀🚀"
exec "$@"       # basically means that the script should run with all variables passed to it