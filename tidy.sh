#!/usr/bin/env bash

set -e
for FILE in $@; do
	echo $FILE | if grep --quiet *.go ; then
		goimports -w $FILE
	fi
done

go mod tidy
