#!/bin/sh

echo "Running linter..."
revive -exclude vendor/... -formatter friendly ./...
status=$?

if [ $status -ne 0 ]; then
  echo "Lint failed"
  exit $status
fi

echo "Running mongodb..."
mongod --dbpath=./data > /dev/null &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start mongodb: $status"
  exit $status
fi

sleep 1

echo "Running tests..."
if ./routes.test; then
  echo "Tests passed!"
  exit 0
else
  echo "Tests failed!"
  exit 1
fi
