# CLI Greeter

A simple, production-ready command-line application that greets users by name.

## Features

- Configurable greeting name via command-line flag
- Environment variable support for default name
- Proper error handling and exit codes
- Comprehensive test coverage

## Installation

### From Source

```bash
make install
```

Or using Go directly:

```bash
go install
```

### Build Binary

```bash
make build
```

The binary will be created in the `bin/` directory.

## Usage

### Basic Usage

```bash
# Use default name "World"
greeter

# Specify name via flag
greeter -name Alice

# Use environment variable
export GREET_NAME=Bob
greeter
```

### Flags

- `-name`: Name to greet (default: "World")

### Environment Variables

- `GREET_NAME`: Default name to use when no flag is provided

### Precedence

The application follows this precedence order (highest to lowest):

1. Command-line flag (`-name`)
2. Environment variable (`GREET_NAME`)
3. Default value (`World`)

## Development

### Prerequisites

- Go 1.25 or later
- golangci-lint (for linting)

### Running Tests

```bash
make test
```

### Running Linter

```bash
make lint
```

### Formatting Code

```bash
make fmt
```

### Running All Checks

```bash
make check
```

This runs formatting, vetting, linting, and tests.

## Project Structure

```
.
├── main.go           # Main application code
├── main_test.go      # Test suite
├── go.mod            # Go module definition
├── makefile          # Build automation
├── .golangci.yml     # Linter configuration
└── README.md         # This file
```

## Exit Codes

- `0`: Success
- `1`: Error occurred

## License

This is a sample project for demonstration purposes.
