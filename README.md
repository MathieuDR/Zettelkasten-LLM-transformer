# Zettelkasten LLM Transformer

Transform personal Zettelkasten notes into a format optimized for Large Language Models.

## Quick Start

```bash
# Using Nix
nix build

# Using Go
go build -o zk-transformer
./zk-transformer -o output_dir
```

## Requirements
- Go 1.23.3+
- Nix (optional)

## Usage
```bash
zk-transformer -o <output_dir> [flags]

Flags:
  -o, --output       Output destination (required)
  -s, --summary      Show processing summary
  -f, --single-file  Single file output (default: true)
```

## Development

This project uses Nix flakes for reproducible builds and development environments:
```bash
# Enter dev environment
nix develop

# Format code (Go/Nix/Markdown)
nix fmt
```

## Building

```bash
# Development build
go build -o zk-transformer

# Nix build with version information
nix build
```
