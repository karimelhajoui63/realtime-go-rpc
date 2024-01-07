# POC realtime in GO/RPC (Connect)

The server send a stream to the clients (like websocket or SSE)

https://github.com/karimelhajoui63/realtime-go-rpc/assets/44633381/e6bdada5-4047-4eba-9f2f-e136cdb11583

Note: After 2 windows (by browser) the calls to the servers are buffered. I don't know why, but it's certainly a browser limitation.

## Quick start

#### With [Task](https://taskfile.dev/) (easier)

Just run:
```sh
task docker:up_n_open  # need: brew install go-task
```

_Tip: you can see all available tasks we the command: `task --list`_

#### ... or with Docker (still easy)

Run:
```sh
docker-compose up
```

then visit: http://localhost:45871

#### ... or even without Docker (can by tricky)

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

## Contribution

You can easly contribute, without having to install anything beside [Docker](https://www.docker.com/), [Docker Desktop](https://www.docker.com/products/docker-desktop/) (personally, I used [OrbStack](https://orbstack.dev/) because it seems faster) and [Visual Studio Code](https://code.visualstudio.com/).

Once the repo is downloaded locally, you can run the command from the VSCode palette: `Dev Containers: Reopen in Container`

Here, you have your environment setup with all you need.
You can browse the [Taskfile](Taskfile.yml) to see useful commands.

_Pro tip: use [Dev Environments](https://chromewebstore.google.com/detail/dev-environments/gnagpachnalcofcblcgdbofnfakdbeka) chrome extension to open the repo in one click from GitHub_


## TODO

 - [ ] Use watermill instead of the rabbitmq lib
 - [ ] Add some tests (just for the XP)
 - [ ] Add impressions on DX
 - [ ] Use 1 connection with the RPC server per sqare on the UI (not only 1 for all of them)
 - [ ] Docker: clean hard coded constants
 - [ ] Docker: generate proto file within the docker build (?)

## Troubleshooting 

 - Cannot `git push` because of `Permission denied (publickey)`?

   - Run `ssh-add $HOME/.ssh/github_rsa` ([source](https://code.visualstudio.com/remote/advancedcontainers/sharing-git-credentials))