# cobraCLI

CLI client for interacting with the distributed-system services in
`goDistributedSystem`.

## Why this repo depends on `goDistributedSystem`

This CLI imports generated protobuf types from `goDistributedSystem/pkg/pb`
(used by `internal/client/grpc.go`).

Because of that, this repo is not fully standalone: you also need the
`goDistributedSystem` repository available locally.

- Required repo: `https://github.com/unsubstantiated-Script/goDistributedSystem`
- Current local module mapping in `go.mod`:

```go
replace goDistributedSystem => ../goDistributedSystem
```

## Setup both repos locally

Clone both repositories as siblings so the `replace` path resolves correctly.

```zsh
cd /home/skripty/GolandProjects
git clone https://github.com/unsubstantiated-Script/goDistributedSystem.git
git clone https://github.com/unsubstantiated-Script/cobraCLI.git
```

Expected layout:

```text
/home/skripty/GolandProjects/
  cobraCLI/
  goDistributedSystem/
```

## Follow server setup from `goDistributedSystem`

After cloning, open the `goDistributedSystem` README and follow its setup/run
instructions first. This CLI assumes those HTTP/gRPC services are available.

Repository: `https://github.com/unsubstantiated-Script/goDistributedSystem`

## Verify CLI module wiring

Run these inside `cobraCLI`:

```zsh
cd /home/skripty/GolandProjects/cobraCLI
go mod tidy
go list -m goDistributedSystem
go test ./...
```

If `goDistributedSystem/pkg/pb` cannot be found, confirm the sibling folder
exists or update the `replace` path in `go.mod` for your local environment.

## Build and run

```zsh
cd /home/skripty/GolandProjects/cobraCLI
go build -o cobraCLI .
./cobraCLI --help
```

## Common commands

Use default addresses unless you need to override them:

- `--http-addr` default: `:8080`
- `--grpc-addr` default: `:50051`

Examples:

```zsh
./cobraCLI status
./cobraCLI submit --command "echo hello"
./cobraCLI workerPing --name local-worker
```

Override endpoints when needed:

```zsh
./cobraCLI status --http-addr http://127.0.0.1:8080
./cobraCLI workerPing --grpc-addr 127.0.0.1:50051 --name worker-a
```

## Troubleshooting

- `missing protocol scheme` for HTTP calls: pass `--http-addr` as
  `http://host:port` if your environment requires an explicit scheme.
- `module goDistributedSystem ... not found`: verify local folder layout or
  adjust the `replace` directive.
- gRPC connection errors: confirm `goDistributedSystem` services are running
  and reachable on the configured ports.
