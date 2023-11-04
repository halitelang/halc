# halc Compiler

Welcome to the official repository of the `halc` compiler, the primary compiler for the Halite programming language. Halite is a general-purpose language that combines the efficiency of compiled languages with a minimal and readable syntax. Designed with modern programming practices in mind, Halite offers an extensive, modular standard library, easy parallelism through Fibers (coroutines), and a prototype-based procedural paradigm.

## Features

- **Compiled to Machine Code**: Enjoy the performance of direct machine-level execution.
- **Prototype-Based Procedural Paradigm**: A flexible approach to procedural programming without the constraints of classes.
- **Modular Standard Library**: Use only what you need, keeping your projects lightweight and efficient.
- **Easy Parallelism**: Halite’s Fibers make writing concurrent code straightforward and maintainable.
- **Minimal Syntax**: Focus on your code logic with a syntax that is easy to write and read.

## Getting Started

### Prerequisites

- **Go**: Halite compiler is built in Go. Ensure you have the latest version of Go installed on your machine.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/halitelang/halc.git
cd halc
```

Build the compiler:

```bash
go build -o halc
```

### Usage

Compile a Halite source file:

```bash
./halc myprogram.hal
```

This will compile `myprogram.hal` to an executable machine code file.

## Documentation

For more details on the Halite language and its features, visit the [Halite Language Documentation](#) (Link to be added).

## Contributing

Contributions to the `halc` compiler are welcome! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to submit pull requests, the process for submitting bugs, and other ways you can contribute to the Halite community.

## Support

If you encounter any issues or have questions, please file an issue on the GitHub issue tracker.

## License

`halc` is distributed under the [MIT License](LICENSE).
