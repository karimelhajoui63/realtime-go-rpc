# POC realtime in GO/RPC (Connect)

The server send a stream to the clients (like websocket or SSE)

https://github.com/karimelhajoui63/realtime-go-rpc/assets/44633381/b93df19f-b373-457e-bf8f-cca2b61c6878

Note: After 2 windows (by browser) the calls to the servers are buffered. I don't know why, but it's certainly a browser limitation.

<br>

## Quick start

### With [Task](https://taskfile.dev/) (easier)

Just run:
```sh
task docker:up_n_open  # need: brew install go-task
```

_Tip: you can see all available tasks we the command: `task --list`_

### ... or with Docker (still easy)

Run:
```sh
docker-compose up
```

then visit: http://localhost:45871

### ... or even without Docker (can by tricky)

In one terminal:
```sh
cd backend
go run cmd/rpcserver_[rabbitmq|watermill]/main.go
```

In an other one:
```sh
cd frontend
npm run dev
```

(See `README.md` in `frontend` and `backend` for more info)

<br>

## Contribution

You can easly contribute, without having to install anything beside [Docker](https://www.docker.com/), [Docker Desktop](https://www.docker.com/products/docker-desktop/) (personally, I used [OrbStack](https://orbstack.dev/) because it seems faster) and [Visual Studio Code](https://code.visualstudio.com/).

Once the repo is downloaded locally, you can run the command from the VSCode palette: `Dev Containers: Reopen in Container`

Here, you have your environment setup with all you need.
You can browse the [Taskfile](Taskfile.yml) to see useful commands.
You can even start docker containers within the dev container (i.e. `docker-compose up rabbitmq` for exemple if you need a rabbitmq instance)

_Pro tip: use [Dev Environments](https://chromewebstore.google.com/detail/dev-environments/gnagpachnalcofcblcgdbofnfakdbeka) chrome extension to open the repo in one click from GitHub_

<br>


## To-do list

 - [ ] Docker: clean hard coded constants
 - [ ] Frontend: use 1 connection with the RPC server per sqare on the UI (not only 1 for all of them)
 - [ ] Refacto: use [Wire](https://github.com/google/wire) to provide cleaner DI (?)

<br>

## Impressions on DX

- Golang

The **typing system is really cool** and bring a lot of feature that improve the DX.
Unlike Python, when I run the program, it works on the first time (because the potentiel runtime bug is catch directly from the IDE wile coding) 🔥

- Hexagonal architecture

It helps to separate logic that souldn't be in the same files, and it really makes it **easier to add/swipe an adapter** 🔥

- Connect RPC

**Way easier** to generate proto stub than with gRPC. It worked the first time (not the case - *at all* - of gRPC).
Plus, it's retro-compatible with the gRPC's API 🔥

- Dev containers

No **more local env issues** (plus, it really smooth, identical to local) 🔥

<br>


## Troubleshooting 

 - Cannot `git push` because of `Permission denied (publickey)`?

   - Run `ssh-add $HOME/.ssh/github_rsa` ([source](https://code.visualstudio.com/remote/advancedcontainers/sharing-git-credentials))
