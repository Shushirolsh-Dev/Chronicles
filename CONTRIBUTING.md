# Contributing

Thanks for considering a contribution to Chronicles. The project values small, focused changes and a calm, dependable user experience.

## Before you start

- Check existing issues and pull requests for related work.
- For substantial changes, open an issue first to discuss the direction.
- Keep the CLI behavior predictable and avoid adding configuration or dependencies without a clear need.

## Development setup

Requirements: Go 1.22+ and Git.

```sh
git clone https://github.com/Shushirolsh-Dev/Chronicles.git
cd Chronicles
go test ./...
go vet ./...
make build
```

## Pull requests

- Keep changes focused and explain the user problem they solve.
- Add or update tests for behavior changes.
- Update README or command documentation when flags or output change.
- Run `gofmt`, `go test ./...`, and `go vet ./...` before opening a PR.
- Do not include generated binaries or unrelated formatting changes.

## Commit messages

Use concise, imperative messages. Examples:

- `feat: add date range filtering`
- `fix: handle empty Git history`
- `docs: clarify changelog output`

## Project principles

- Local-first and offline by default.
- Read-only interaction with the user's repository.
- Clear errors and useful defaults.
- Minimal dependencies and idiomatic Go.
- No telemetry, hidden network calls, or unnecessary interface noise.
