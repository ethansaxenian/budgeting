#!/bin/sh -e

BACKUP_FILE=backups/backup_$(date +"%Y-%m-%dT%H:%M:%S%z").sql
echo "Creating $BACKUP_FILE..."
pg_dump --data-only --exclude-table=goose_db_version "$PSQL_CONNECTION_STRING" > "$BACKUP_FILE"
sleep 86400
