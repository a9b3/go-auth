#!/usr/bin/env bash
# Run this script and supply it with a valid .env file
#
# ex.
# ./scripts/migrate -c .env

# Require goose
if ! [ -x "$(command -v goose)" ]; then
  echo 'Error: goose is not installed.' >&2
  exit 1
fi

# Set variables based on supplied flags
while getopts 'c:' flag; do
  case "${flag}" in
    c) config="${OPTARG}" ;;
  esac
done

# Error checking the supplied flags
if [ ! -f "$config" ]; then
  echo 'Error: must provide valid file to -c flag.' >&2
  exit 1
fi

source "$config"

goose -dir migrations postgres \
  "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_SSLMODE"\
  up
