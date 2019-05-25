#!/usr/bin/bash

rm -rf ./models && sqlboiler --config ste_files/sqlboiler.toml psql && go test ./models
