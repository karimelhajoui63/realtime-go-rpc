[Official documentation](https://connectrpc.com/docs/web/getting-started) of Connect


### Commands used:

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
