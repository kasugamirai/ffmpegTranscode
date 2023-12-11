#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=videoHost
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}