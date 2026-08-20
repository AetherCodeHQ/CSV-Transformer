# CSV Transformer

![CI](https://github.com/Qyroxen/CSV-Transformer/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/CSV-Transformer?style=social)

> A powerful CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/CSV-Transformer?style=social)](https://github.com/Qyroxen/CSV-Transformer/stargazers)

## What is it?

CSV Transformer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/CSV-Transformer.git
cd CSV-Transformer
go build -o csvtransformer .

# Run
./csvtransformer --help
```

## CLI Usage

```bash
# Basic usage
./csvtransformer

# With flags
./csvtransformer --verbose --output json

# Get help
./csvtransformer --help
```

## Examples

```bash
# Example 1
./csvtransformer example1

# Example 2
./csvtransformer example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o csvtransformer .

# Lint
go vet ./...
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/CSV-Transformer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/CSV-Transformer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/CSV-Transformer/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/CSV-Transformer?style=social" alt="Fork this repo">
  </a>
</p>
