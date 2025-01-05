#!/usr/bin/env bash
ls
unset dirs files
dirs=$(go list -f {{.Dir}} ./... | grep -v /yotiprotoshare/ | grep -v /yotiprotocom/ | grep -v /yotiprotoattr/)
for d in $dirs; do
  for f in $d/*.go; do
    files="${files} $f"
  done
done
if [ -n "$(gofmt -d $files)" ]; then
  echo "Go code is not formatted:"
  gofmt -d .
  exit 1
fi
