# Contributing to halc

We are excited that you are interested in contributing to halc! This document provides guidelines for contributions to the Halite language compiler. Please read through this document before submitting your contributions to ensure a smooth and efficient workflow for everyone.

## How to Contribute

### Reporting Bugs

- Use the GitHub issues tracker to report bugs.
- Before you file an issue, please search the existing issues to avoid duplicates.
- When creating a bug report, include a clear title and description, as much relevant information as possible, and a code sample if applicable.

### Suggesting Enhancements

- For feature requests, use the GitHub issues tracker as well.
- Provide a clear title and description of the enhancement.
- Explain why this enhancement would be useful to most Halite users.

### Pull Requests

- Fork the repository and create your branch from `master`.
- If you've added code that should be tested, add tests.
- Ensure your code follows the existing code style.
- Write meaningful commit messages.
- Include appropriate documentation in the `README.md`.
- Submit your pull request with a clear title and description.

## Development Setup

1. Fork the repo and clone it to your development machine.
2. Ensure you have a working Go environment.
3. Set up your local development environment:
   ```bash
   go mod download
   go build
   ```
4. Make sure the tests pass on your local machine:
   ```bash
   go test ./...
   ```

## Code Review Process

The core team looks at Pull Requests on a regular basis. After feedback has been given, we expect responses within two weeks. After two weeks, we may close the PR if it isn't showing any activity.

## Coding Conventions

- Use gofmt for formatting your code.
- Keep your code idiomatic to Go.
- Follow the naming conventions for files, variables, and functions as seen in the current codebase.
- Comment your code where necessary but strive for self-explanatory code.

## License

By contributing to halc, you agree that your contributions will be licensed under the same MIT License that covers the project.

Thank you for your contribution!
