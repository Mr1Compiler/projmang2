#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./check_health.sh <BACKEND_URL>"
  echo "Example: ./check_health.sh https://backend-production.up.railway.app"
  exit 1
fi

URL="$1/health"
echo "Checking health of $URL..."

RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

if [ "$RESPONSE" == "200" ]; then
  echo "✅ Backend is healthy (Status 200)"
  curl -s "$URL"
  echo ""
else
  echo "❌ Backend check failed (Status $RESPONSE)"
  echo "Response body:"
  curl -s "$URL"
  echo ""
fi
