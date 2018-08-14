# Go Auth

Auth service

## Prereqs

- [dep](https://github.com/golang/dep): dependencies
- [canthefason/go-watcher](https://github.com/canthefason/go-watcher): automatically restarting server upon file change during dev
- [pressly/goose](https://github.com/pressly/goose): working with migrations

## Debugging

### PostGres

You can either use a gui like [tableplus](https://tableplus.io/) or psql
`docker exec -it auth_postgres_1 psql -U postgres -d postgres` (you have to use
docker exec because native psql depends on having some postgres files on your host
machine).
