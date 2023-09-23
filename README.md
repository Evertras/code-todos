# code-todos

Programmatically find and track TODO statements in codebases.

## Generated example

See [./todos.md](./todos.md) for an example of a generated markdown table.

## Supported languages

For now, just Go.  May add others later if useful.

## Usage

```bash
# Find TODOs in current directory recursively
code-todos find .

# Find TODOs in pkg and cmd dirs only
code-todos find pkg cmd

# Find TODOs in specific files
code-todos find main.go pkg/thing/library.go
```
