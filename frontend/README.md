[Official documentation](https://connectrpc.com/docs/web/getting-started) of Connect


### Commands used in bulk (just for history purpose, do not use them directly):

```sh
npm create vite@latest frontend -- --template react-ts
cd frontend
npm install

npm install --save-dev @bufbuild/buf @connectrpc/protoc-gen-connect-es @bufbuild/protoc-gen-es
npm install @connectrpc/connect @connectrpc/connect-web @bufbuild/protobuf

npx buf lint ../proto
npx buf generate --template ../buf.gen.ts.yaml ../proto

npm run dev
```

### Bref explanation

In the `frontend` directory, there is just a node server that serve a page that interact via RPC the server.
The main objectif is to see in action the streaming communication `server -> client`.

Warning: only 2 windows seem to be able to communicate with the server simultaneously.

(I could have done the same with Go, but it was faster for me to do with NodeJS)

### Core architecture

Nothing special, I just ran `npm create vite ...` since the client isn't really important