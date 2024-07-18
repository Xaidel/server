#!/usr/bin/env sh

mysql -u root -e "CREATE DATABASE IF NOT EXISTS csprobe_go"
echo "Connecting to CSPROBE_GO database"

go run ./migrations/migrate.go

go run main.go