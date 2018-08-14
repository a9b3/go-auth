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

while getopts 'c:' flag; do
  case "${flag}" in
    c) config="${OPTARG}" ;;
  esac
done

if [ ! -f "$config" ]; then
  echo 'Error: must provide valid file to -c flag.' >&2
  exit 1
fi

source "$config"

goose -dir migrations postgres \
  "postgres://$POSTGRESUSER:$POSTGRESPASSWORD@$POSTGRESHOST:$POSTGRESPORT/$POSTGRESDB?sslmode=disable"\
  up
