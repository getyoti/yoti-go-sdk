#!/usr/bin/env bash

set -e
goimports -w .
go mod tidy
