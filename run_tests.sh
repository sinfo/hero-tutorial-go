#!/bin/sh

# Start the first process
mongod --dbpath=./data &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start mongodb: $status"
  exit $status
fi

sleep 1

if ./routes.test; then
  echo "Tests passed!"
  exit 0
else
  echo "Tests failed!"
  exit 1
fi
