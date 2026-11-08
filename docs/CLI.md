# CLI reference

Chronicles has two commands: the default journal and `changelog`.

## `chronicle`

Displays commits from the last 14 days, grouped by the author's commit date.

```sh
chronicle [flags]
```

## `chronicle changelog`

Writes a Markdown changelog to standard output.

```sh
chronicle changelog [flags]
```

## Shared flags

| Flag | Default | Meaning |
| --- | --- | --- |
| `--days` | `14` | Look back this many days. Values below 1 are rejected. |
| `--path` | `.` | Repository directory. A nested directory inside a repository is accepted. |
| `--author` | empty | Git author pattern, matched against author name or email. |

## Examples

```sh
# Current repository, default 14-day window
chronicle

# Last week
chronicle --days 7

# Filter by author
chronicle --author "Ada"

# Another local repository
chronicle --path ../service

# Save Markdown output
chronicle changelog --days 45 > CHANGELOG.md
```

## Exit behavior

- Successful execution exits with status 0.
- Invalid paths, non-repository paths, invalid day counts, and Git execution errors are reported to standard error and exit non-zero.
- An empty history is a successful result and produces an explanatory message.
