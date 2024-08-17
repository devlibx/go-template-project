#!/bin/bash

# Check if the file exists
if [ ! -f "./env/common.env" ]; then
    echo "Error: common.env not found!"
    exit 1
fi

# Read the file line by line
while IFS= read -r line; do
  # Skip empty lines
  if [ -z "$line" ]; then
      continue
  fi

  # Export each line as an environment variable
  export "$line"
done < "./env/common.env"

# Check if the file exists
if [ ! -f "./env/stage.env" ]; then
    echo "Error: stage.env not found!"
    exit 1
fi

# Read the file line by line
while IFS= read -r line; do
  # Skip empty lines
  if [ -z "$line" ]; then
      continue
  fi

  # Export each line as an environment variable
  export "$line"
done < "./env/stage.env"

go run cmd/server/main.go