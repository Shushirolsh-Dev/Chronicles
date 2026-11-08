# Chronicles

**A calm, readable development journal generated from your Git history.**

Chronicles turns commits in your local Git repositories into a human-readable development journal or Markdown changelog. It works offline, needs no account or configuration, and leaves your repository untouched.

[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Why Chronicles?

Git records what changed. Chronicles makes that history easier to revisit.

Instead of scanning a dense commit log, get a dated overview of your work. Use it for a daily review, a weekly recap, or release notes you can share.

- **Local-first:** reads your local Git repository; no network access or account required.
- **Readable:** groups commits by date with a clear visual hierarchy.
- **Useful filters:** narrow results by time period, repository path, or author.
- **Markdown output:** generate changelogs for READMEs, release notes, and documentation.
- **Minimal:** focused Go CLI, sensible defaults, and a small dependency set.

## Requirements

- Go 1.22 or newer
- Git installed and available on your `PATH`
- A local Git repository containing commit history

## Installation

### Install with Go

```sh
go install github.com/Shushirolsh-Dev/Chronicles/cmd/chronicle@latest
```

### Build from source

```sh
git clone https://github.com/Shushirolsh-Dev/Chronicles.git
cd Chronicles
make build
```

The executable is written to `./bin/chronicle`. To install it into your Go binary directory, run `make install`.

## Quick start

Run Chronicles from inside any Git repository:

```sh
chronicle
```

Or point it at a repository explicitly:

```sh
chronicle --path ~/code/my-project
```

By default, Chronicles shows commits from the last 14 days.

## Commands

### Development journal

```sh
chronicle
chronicle --days 7
chronicle --days 30 --author "Jonathan"
chronicle --path ~/code/api --days 14
```

The journal groups commits by the author's commit date. Each entry includes the abbreviated hash, subject, and author.

### Markdown changelog

```sh
chronicle changelog
chronicle changelog --days 30
chronicle changelog --path ~/code/api --author "Jonathan"
chronicle changelog --days 90 > CHANGELOG.md
```

Changelog output goes to standard output, so it can be redirected or piped into another tool.

## Options

| Flag | Default | Description |
| --- | --- | --- |
| `--days` | `14` | Include commits from the last N days. Must be greater than zero. |
| `--path` | `.` | Path to the Git repository to inspect. |
| `--author` | — | Filter by author name or email (case-insensitive Git author match). |
| `--help` | — | Show command help. |

Options work with both the default journal command and `changelog`.

## Example output

Dates and commit messages below are illustrative; actual output comes from your repository.

```text
  CHRONICLES
  my-project  ·  last 14 days

  TODAY · 08 NOV 2026
  ──────────────────────────────────────────────
  8f31c2a  Add request validation
           Jonathan

  2d0a1be  Simplify configuration loading
           Jonathan

  YESTERDAY · 07 NOV 2026
  ──────────────────────────────────────────────
  6a9d20f  Document deployment workflow
           Jonathan

  3 commits across 2 days
```

## How it works

Chronicles invokes Git's read-only `git log` command against the selected repository. It reads commit hashes, author details, author timestamps, subjects, and bodies, then groups and renders the results.

- It does not checkout branches, modify files, create commits, or change Git configuration.
- It does not contact GitHub or any remote service.
- It uses author dates for journal grouping and the `--since` cutoff.
- Merge commits are included when Git reports them in the selected history.
- Commits without a subject receive a safe fallback label.

## Development

```sh
git clone https://github.com/Shushirolsh-Dev/Chronicles.git
cd Chronicles
go test ./...
go vet ./...
make build
```

### Project layout

```text
.
├── cmd/chronicle/main.go
├── internal/
│   ├── git/repository.go
│   ├── journal/journal.go
│   ├── journal/journal_test.go
│   └── render/render.go
├── .gitignore
├── go.mod
├── LICENSE
├── Makefile
└── README.md
```

## Design principles

Chronicles is intentionally restrained. It favors clear hierarchy, useful defaults, predictable output, and a small surface area over dashboards, telemetry, configuration files, or decorative noise.

## License

Released under the [MIT License](LICENSE).
