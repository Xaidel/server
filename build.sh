#!/usr/bin/env sh
go run ./migrations/migrate.go

go run main.go