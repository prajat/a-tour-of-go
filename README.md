# A Tour of Go

Practicing the official "A Tour of Go" example programs for revision and hands-on learning.

## Contents

- Purpose
- What’s included
- Prerequisites
- Quick start
- Project structure
- Notes & tips
- Editor & Git
- Resources
- License

## Purpose

This repository contains small example programs and exercises based on the official "A Tour of Go" (https://go.dev/tour/list). It’s intended for quick revision, practice, and reference when learning or refreshing core Go concepts.

## What’s included

- Short, focused Go files that map to lessons/exercises from the Tour of Go
- Minimal, easy-to-read examples suitable for quick review and modification
- Links to the original Tour of Go pages for deeper study

> Adjust or extend the examples to match your personal study plan.

## Prerequisites

- Go installed (recommended: latest stable release)  
  Install: https://go.dev/doc/install
- Basic familiarity with the command line
- GOPATH / module-aware workflow (Go modules are recommended)

## Quick start

1. Clone the repository:
   1. `git clone <repo-url>`
2. Change into the folder for the example you want to run:
   1. `cd examples/<example-folder>`
3. Run the example:
   1. `go run .` (or `go run filename.go`)

For running and testing multiple small files, you can use `go run` or build binaries with `go build`.

## Project structure (example)

| Path              | Description                                         |
| ----------------- | --------------------------------------------------- |
| basics/           | Example programs organized by Tour chapter or topic |
| basics/variables/ | Simple variable and type examples                   |
| basics/control/   | If/for/switch examples                              |
| basics/functions/ | Function and closure examples                       |
| README.md         | This file with project overview and instructions    |

Adjust the structure to match the content you add.

## Notes & tips

- Use Go modules for dependency management: `go mod init <module-name>`
- Keep examples small and focused — one concept per file makes revision faster.
- Add comments and links to the corresponding Tour of Go pages for future reference.
- Run `gofmt -w .` to keep formatting consistent.
- Use `go vet` and `golint` (or `staticcheck`) to catch issues and improve code quality.

## Editor & Git

- Editor: I use Zed for editing — a lightweight, fast editor with modern ergonomics. Configure Zed with Go extensions (language server) for auto-completion and formatting.
- Git: I use lazygit for interactive, terminal-based Git workflow. It's useful for staging, committing, rebasing, and reviewing diffs quickly.
- Recommended workflow:
  1. Edit code in Zed with gopls (Go language server) enabled.
  2. Run `gofmt -w .` or configure format-on-save.
  3. Use lazygit to stage and commit changes, and to push branches.

## Resources

- A Tour of Go — https://go.dev/tour/list
- Official Go documentation — https://go.dev/doc/
- Go modules — https://go.dev/ref/mod
- Zed editor — https://zed.dev
- lazygit — https://github.com/jesseduffield/lazygit

## License

Specify a license for your repository (MIT, Apache-2.0, etc.). Example:

MIT License — see LICENSE file for details.

---

If you want, I can:

- Add Zed configuration snippets (settings, recommended extensions).
- Add a sample lazygit config and common commands.
- Insert example filenames and short code snippets from this repo.
