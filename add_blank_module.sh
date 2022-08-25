#!/usr/bin/env bash

moduleName=$1
echo "Create directory for module \"$moduleName\""
mkdir ./internal/modules/$moduleName

for i in dto entities repositories services; do
    echo "Create directory \"$i\" for module \"$moduleName\""
    mkdir ./internal/modules/$moduleName/$i

    echo "Create empty file for module $moduleName/$i"
    echo "package $i" > internal/modules/$moduleName/$i/$moduleName.go
done