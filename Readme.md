# Ecommerce

## Frontend
SolidJS, Tailwind

## Backend
Go-Echo, MongoDB, Postgres, ProtoBuffs

## Major Backend TODOs:
- [ ] Add ProtoBuffer
- [ ] Dockerize
- [ ] Authentication
- [ ] Caching
- [ ] Modularize Go Backend
- [ ] Make Backend use Concurrency
- [ ] ENV Variables

## Updating proto buffers

```bash
protoc -I=$PWD/protobuffers --go_out=$PWD/server --ts_out=$PWD/client/src $PWD/protobuffers/*.proto
```

## Scripts

```bash
# Run the server
nodemon --exec go run *.go --signal SIGTERM

# Run the client
bun run dev
```