# Ecommerce

## Frontend
SolidJS, Tailwind

## Backend
- Go-Echo - Framework,
- MongoDB - get details of product/user
- Postgres - Primary Datastore
- Go-Orm
- ProtoBuffs - For server to client communication. (Not sure whether to use it or not)

## Major Backend TODOs:
- [ ] Dockerize
- [ ] Authentication
- [ ] Caching
- [ ] Modularize Go Backend
- [ ] Make Backend use Concurrency
- [ ] ENV Variables
- [ ] meilisearch

## Updating proto buffers

```bash
protoc -I=$PWD/protobuffers --go_out=$PWD/server --ts_out=$PWD/client/src $PWD/protobuffers/*.proto
```

## Scripts

```bash
# Run the server
go mod download
nodemon --exec go run *.go --signal SIGTERM

# Run the client
bun install
bun run dev
```