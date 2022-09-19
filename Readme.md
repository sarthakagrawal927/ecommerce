# Ecommerce

## Frontend
SolidJS, Tailwind, Typescript

## Backend
Go-Echo, MongoDB, Postgres, ProtoBuffs

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