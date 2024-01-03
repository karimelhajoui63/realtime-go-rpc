# POC realtime in GO/RPC (Connect)

The server send a stream to the clients (like websocket or SSE)

https://github.com/karimelhajoui63/realtime-go-rpc/assets/44633381/e6bdada5-4047-4eba-9f2f-e136cdb11583

Note: After 2 windows (by browser) the calls to the servers are buffered. I don't know why, but it's certainly a browser limitation.

### Quick start

In one terminal:
```sh
cd backend
go run cmd/rpcserver/main.go
```

In an other one:
```sh
cd frontend
npm run dev
```

(See `README.md` in `frontend` and `backend` for more info)

### TODO 

 - [ ] Add impressions on DX
 - [ ] Use watermill instead of the rabbitmq lib
 - [ ] Add docker-compose
 - [ ] Add gofumpt
 - [ ] Add dev environment w/ docker
 - [ ] Add Taskfile?
 - [ ] Add some tests (just for the XP)