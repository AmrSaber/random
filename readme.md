# Random Generator CLI

> Old package published on npm has been removed since it was broken and I don't intend to fix it.

Generates and prints a secure random string (and other types) to terminal.

## Install

### Using Homebrew

```bash
brew tap AmrSaber/tap
brew install random
```

### Using Go

Install `go` then run

```bash
go install github.com/AmrSaber/random/v3@latest
```

### Using mise

To install with [mise](https://mise.jdx.dev). First enable experimental features, since Go backend is still experimental.

```bash
mise settings experimental=true
```

The install with

```bash
mise use -g go:github.com/AmrSaber/random/v3@latest
```

Using mise makes it easier to stay updated with the latest versions, you typically update installed packages with `mise up`.

## Usage

However way you use to install the package, it will be available in the CLI as `random`.

After installation, use `random <command>`, you can use `random <command> -h` to show help message related to that command, or use `random -h` to list all the available commands and options.

### Commands

Available commands are [`string`, `pick`, `shuffle`, `id`].

## Feature requests and bug reports

If you have a bug or have a good idea for a new feature, please [open an issue](https://github.com/AmrSaber/rand-string/issues).
