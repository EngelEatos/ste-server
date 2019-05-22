#!/usr/bin/bash

rm -rf ./models && sqlboiler psql && go test ./models
