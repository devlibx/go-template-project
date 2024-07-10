#!/bin/bash

# Check if the file exists
if [ ! -f "common.env" ]; then
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
done < "common.env"

# Check if the file exists
if [ ! -f "stage.env" ]; then
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
done < "stage.env"

go run cmd/server/main.go