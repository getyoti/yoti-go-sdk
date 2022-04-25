#!/usr/bin/env bash
go build ./...

for d in _examples/*/; do
  (cd "$d" && go mod tidy -compat=1.17)
done
