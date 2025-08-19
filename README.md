# Zap Fluent Encoder

A fluent interface encoder for Uber's Zap logging library that provides a more intuitive way to add structured logging
fields.

## ⚠️ Work in Progress

This project is currently under active development. APIs and functionality may change without notice.

## Development

This library is being developed with the assistance of Jules by Google, an AI-powered software engineer. The following rules are being followed during development:

These rules serve as a strict guideline to ensure code quality, readability, and maintainability. However, they are not unbreakable laws; exceptions can be made when adhering to a rule would detract from the overall quality of the code.

- **Prefer Self-Documenting Code**: Avoid comments by using descriptive function and variable names that make the code's purpose clear. Comments should only be used when the code's logic is inherently complex and cannot be clarified through refactoring.
- **AAA Unit Test Structure**: All unit tests must be implicitly structured following the Arrange-Act-Assert (AAA) pattern.
- **Complete Test Coverage**: All production code must be covered by at least one unit test.
- **Standard Import Formatting**: All import blocks must be grouped into four categories in a specific order: standard library, third-party, shared internal modules, and intra-module dependencies.
- **No Unnecessary Pointers**: Avoid using pointers for optional values or to pass arguments to functions. Pointers should only be used when strictly necessary (e.g., for channels, or for large objects where copying is a performance concern). Prefer using an `Optional` monad for optional values.
- **SOLID Principles**: Code should adhere to SOLID principles, with a strong emphasis on the Single Responsibility Principle.
- **Max 2 Parameters**: Functions and methods should have at most two parameters. If a third is required, it should typically be a `context.Context`. More complex arguments should be grouped into a struct.
