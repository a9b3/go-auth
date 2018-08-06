# Go Starter

Go server boilerplate.

## Prereqs

- Install [go-watcher](https://github.com/canthefason/go-watcher) for automatically restarting server upon file change during dev.
- Use [dep](https://github.com/golang/dep) for dependencies.

## Debugging

### PostGres

You can either use a gui like [tableplus](https://tableplus.io/) or psql
`docker exec -it auth_postgres_1 psql -U postgres -d postgres` (you have to use
docker exec because native psql depends on having some postgres files on your host
machine).
