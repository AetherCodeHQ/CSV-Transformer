# CSV Transformer

![CI](https://github.com/Qyroxen/CSV-Transformer/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/CSV-Transformer/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/CSV-Transformer?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/CSV-Transformer)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/CSV-Transformer)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/CSV-Transformer?style=social)](https://github.com/Qyroxen/CSV-Transformer/stargazers)

## What is it?

CSV Transformer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
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
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

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
  <a href="https://github.com/Qyroxen/CSV-Transformer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/CSV-Transformer" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/CSV-Transformer/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/CSV-Transformer" alt="Pull Requests">
  </a>
</p>
