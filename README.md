# code-todos

[![Coverage Status](https://coveralls.io/repos/github/Evertras/code-todos/badge.svg?branch=main)](https://coveralls.io/github/Evertras/code-todos?branch=main)

Programmatically find and track TODO statements in codebases.

## What it does

Uses an AST to find all comments that contain `TODO:` and generate some useful
data about where they are and what they say.  This can then be used either
in further scripting with JSON output, or as a simple Markdown document that
can be generated and updated in a repository.

## Installation

Releases are available as raw binaries on the
[releases page](https://github.com/Evertras/code-todos/releases).

You can also install via Go:

```bash
# Install the binary to ~/some/dir... or omit GOBIN to install to default location
GOBIN=~/some/dir go install github.com/evertras/code-todos/cmd/code-todos@v0.1.0
```

## Generated example

See [./todos.md](./todos.md) for the most up to date example that is generated
for this repository.  A snapshot is shown below for quick reference, but this
may be out of date.

| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
| [internal/todos/find.go](./internal/todos/find.go#L27) | todos | 27 | TODO: add globbing |
| [internal/todos/find.go](./internal/todos/find.go#L48) | todos | 48 | TODO: Filter better |
| [internal/todos/find.go](./internal/todos/find.go#L56) | todos | 56 | TODO: Better error handling |
| [internal/todos/testdata/go/main.go](./internal/todos/testdata/go/main.go#L4) | main | 4 | TODO: Write a really cool thing here. And do it on multiple lines |


## Supported languages

For now, just Go.  May add others later if useful.

A general "catch-all" regex mode may also be helpful in the future.

## Usage

By default, `code-todos` will output a Markdown table of all TODOs it finds.
You can supply `-o json` to output JSON which can be used in further scripting.

The general idea is to keep a file of current TODOs, particularly so that in
a PR you can track added/completed TODOs in the codebase as a health metric.

```bash
# Check help at any time to see available commands and options
code-todos --help

# Find TODOs in current directory recursively
code-todos find .

# Find TODOs in pkg and cmd dirs only
code-todos find pkg cmd

# Find TODOs in specific files and output in JSON for
# further machine parsing elsewhere, scripting, etc.
code-todos find main.go pkg/thing/library.go -o json

# Count TODOs, as an example of combining with jq
code-todos find . -o json | jq length
```
