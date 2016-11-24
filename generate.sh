#!/usr/bin/env bash

go run gen/main.go -o num32 -s 32 -t int32 base
go run gen/main.go -o num64 -s 64 -t int64 base
go run gen/main.go -o num32f -s 32f -t float32 base
go run gen/main.go -o num64f -s 64f -t float64 base