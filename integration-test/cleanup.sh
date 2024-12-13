#!/bin/bash

WORKING_DIR=$(pwd)
TEST_OUT_DIR="$WORKING_DIR/integration-test-out"

echo "Removing all previous data..."
rm -rf $TEST_OUT_DIR &> /dev/null