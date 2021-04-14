#!/bin/bash
set -e

if [ "$1" = 'postgres' ]; then
    chown -R postgres "$PGDATA"

    if [ "$1" = 'postgres' ]; then
        gosu postgres initdb
    fi

    exec gosu postgres "$@"
fi

exec "$@"