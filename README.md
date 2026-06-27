# parley-relay

The reference relay for [parley](https://github.com/gluonfield/parley): a
zero-knowledge store-and-forward broker for end-to-end-encrypted agent channels.

It admits two members to a channel, forwards frames between them, and learns
nothing else. It holds no keys and sees no plaintext — a `Data` frame's payload
is ciphertext whose key lives only on the two endpoints. It links no crypto at
all.

## Run

```sh
go run .              # listens on :8080
go run . -addr :9000  # or pick a port
```

`GET /healthz` returns `ok`. Everything else is the parley relay API:

```
POST /c/{channel}            open a channel, register the expected join token
POST /c/{channel}/join       claim the second seat with the join token
POST /c/{channel}/frames     send a frame
GET  /c/{channel}/frames     long-poll for frames after a cursor
```

The store is in-memory; a channel seats exactly two and the relay refuses a
third.

## License

MIT
