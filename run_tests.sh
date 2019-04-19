#!/bin/sh

echo DB URL $GO_TUTORIAL_DB_URL
sleep 5

if ./routes.test; then
  echo "Tests passed!"
  exit 0
else
  echo "Tests failed!"
  exit 1
fi
