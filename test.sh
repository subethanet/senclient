#!/usr/bin/env bash

for D in `find . -maxdepth 1 -not -path ./.git -not -path ./.idea -type d`
do
    echo "Testing $D"
    go test $D
done

