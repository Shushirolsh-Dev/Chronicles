# Security Policy

## Supported versions

The latest version on the default branch is the supported development version.

## Reporting a vulnerability

Please do not report security vulnerabilities in public issues.

Use GitHub's private vulnerability reporting feature for this repository, if enabled. Include the affected version or commit, reproduction steps, and potential impact. Avoid including real credentials, private repository data, or other sensitive information.

## Security design

Chronicles is designed to inspect local Git history using read-only Git commands. It does not require credentials, make network requests, or modify repository contents. Treat commit messages and author metadata as untrusted text when incorporating them into other systems.
