#!/bin/bash

# Check if host:port and command are provided as arguments
if [ "$#" -lt 3 ]; then
    echo "Usage: $0 <full_url> -- <command>"
    exit 1
fi

FULL_URL=$1
shift  # Remove the full_url argument from the list
shift  # Remove the "--" separator

# Join the remaining arguments to form the command
COMMAND="$*"

S=0
TIMEOUT=30  # Set the timeout value in seconds
ELAPSED=0

while [ $S -ne 200 ] && [ $ELAPSED -lt $TIMEOUT ]
do
    S=$(curl -s -o /dev/null -w "%{http_code}" "$FULL_URL")
    sleep 1
    ELAPSED=$((ELAPSED + 1))
done

if [ $S -ne 200 ]; then
    echo "ERROR: Timeout reached. The API did not start within $TIMEOUT seconds."
    exit 1
fi

echo "The API is UP and RUNNING! ✅"

# Run the provided command
eval "$COMMAND"