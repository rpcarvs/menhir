# apena

`apena` is a personal CLI tool to improve productivity when starting new projects.

It scaffolds simple project structures for:
- Rust
- Python
- Go

## Requirements

This tool runs external commands, so you need these installed:
- Go (to build and run `apena`)
- [Cargo](https://github.com/rust-lang/cargo) (for Rust project creation)
- [uv](https://github.com/astral-sh/uv) (for Python project creation)
- [just](https://github.com/casey/just) (optional, to use the generated `justfile` recipes)

## Install

`apena` can be installed to your GOPATH:

```sh
go build .
```

## Usage

Run the it with one of the language subcommands.

### Go project

To scaffold a Go project with a minimalist dir structure:

```sh
apena go my-go-app
```

### Rust project

To scaffold a Rust project (binary or library) following `cargo init` standards:

```sh
apena rust my-rust-app
apena rust my-rust-lib --lib
```


### Python project

To scaffold a Python project (app, library or package) following `uv init` standards:

```sh
apena python my-python-app
apena python my-python-lib --lib
apena python my-python-package --package
```

