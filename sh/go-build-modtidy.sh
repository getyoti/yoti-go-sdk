#!/usr/bin/env bash
go build ./...

for d in _examples/*/
do
     (cd "$d" && go mod tidy)
done
